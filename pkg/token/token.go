package token

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/taskmeter/config"
	"github.com/hzxiao/taskmeter/pkg/timeutil"
	"strings"
)

var (
	ErrEmptyToken = errors.New("token: empty token")
)

type Context struct {
	ID       string
	Username string
}

func ParseRequest(c *gin.Context) (*Context, error) {
	token, err := getToken(c)
	if err != nil {
		return nil, err
	}
	return Parse(token, config.GetString("jwt_secret"))
}

func getToken(c *gin.Context) (string, error) {
	var token string
	var err error
	methods := strings.Split(config.GetString("token_lookup"), ",")
	for _, method := range methods {
		if token != "" {
			break
		}
		parts := strings.Split(strings.TrimSpace(method), "-")
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		switch k {
		case "header":
			token, err = jwtFromHeader(c, v)
		case "query":
			token, err = jwtFromQuery(c, v)
		case "cookie":
			token, err = jwtFromCookie(c, v)
		}
	}
	if err != nil {
		return "", err
	}
	return token, nil
}

func jwtFromHeader(c *gin.Context, key string) (string, error) {
	auth := c.Request.Header.Get(key)
	if auth == "" {
		return "", ErrEmptyToken
	}

	var t string
	// Parse the header to get the token part.
	fmt.Sscanf(auth, "Bearer %s", &t)
	if t == "" {
		return "", ErrEmptyToken
	}
	return t, nil
}

func jwtFromQuery(c *gin.Context, key string) (string, error) {
	token := c.Query(key)
	if token == "" {
		return "", ErrEmptyToken
	}

	return token, nil
}

func jwtFromCookie(c *gin.Context, key string) (string, error) {
	token, _ := c.Cookie(key)
	if token == "" {
		return "", ErrEmptyToken
	}

	return token, nil
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = claims["id"].(string)
		ctx.Username = claims["username"].(string)
		return ctx, nil
	} else {
		return ctx, err
	}
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

func GenerateToken(ctx Context, secret string) (tokenString string, err error) {
	if secret == "" {
		secret = config.GetString("jwt_secret")
	}
	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       ctx.ID,
		"username": ctx.Username,
		"nbf":      timeutil.Now(),
		"iat":      timeutil.Now(),
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}
