package handler

import (
	"github.com/hzxiao/goutil/assert"
	"testing"
	"github.com/hzxiao/goutil"
)

func TestLogin(t *testing.T) {
	removeAll()

	//sign up first
	_, err := DoSignUp(goutil.Map{"username": "x", "password": "1"})
	assert.NoError(t, err)

	res, err := DoLogin("x", "1")
	assert.NoError(t, err)

	assert.NotNil(t, res.GetString("token"))
}

func TestSignUp(t *testing.T) {
	removeAll()

	result, err := DoSignUp(goutil.Map{"username": "x", "password": "1"})
	assert.NoError(t, err)

	assert.NotNil(t, result)
	assert.Equal(t, "x", result.GetString("username"))

	result, err = DoSignUp(goutil.Map{"username": "x", "password": "1"})
	assert.Error(t, err)
}