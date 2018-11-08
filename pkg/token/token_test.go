package token

import (
	"github.com/hzxiao/goutil/assert"
	"github.com/hzxiao/taskmeter/pkg/timeutil"
	"strconv"
	"testing"
)

func TestToken(t *testing.T) {
	c := Context{
		ID:           "xx",
		Username:     "name",
		GenerateTime: strconv.FormatInt(timeutil.Now(), 10),
		Source:       "Web",
	}
	secret := "secret"
	tokenString, err := GenerateToken(c, secret)
	assert.NoError(t, err)

	parsedCtx, err := Parse(tokenString, secret)
	assert.NoError(t, err)

	assert.Equal(t, c.ID, parsedCtx.ID)
	assert.Equal(t, c.Username, parsedCtx.Username)
	assert.Equal(t, c.GenerateTime, parsedCtx.GenerateTime)
	assert.Equal(t, c.Source, parsedCtx.Source)
}
