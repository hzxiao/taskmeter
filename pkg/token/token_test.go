package token

import (
	"github.com/hzxiao/goutil/assert"
	"testing"
)

func TestToken(t *testing.T) {
	c := Context{
		ID:       "xx",
		Username: "name",
	}
	secret := "secret"
	tokenString, err := GenerateToken(c, secret)
	assert.NoError(t, err)

	parsedCtx, err := Parse(tokenString, secret)
	assert.NoError(t, err)

	assert.Equal(t, c.ID, parsedCtx.ID)
	assert.Equal(t, c.Username, parsedCtx.Username)
}
