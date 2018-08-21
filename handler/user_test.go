package handler

import (
	"github.com/hzxiao/goutil/assert"
	"github.com/hzxiao/taskmeter/pkg/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	res, err := httptest.Get("/")
	assert.NoError(t, err)

	assert.Equal(t, "hello", res.GetStringP("data"))
}
