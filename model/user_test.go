package model

import (
	"testing"
	"github.com/hzxiao/goutil/assert"
	"github.com/hzxiao/goutil"
)

func TestInsertUser(t *testing.T) {
	removeAll()

	user1, err := InsertUser(nil)
	assert.Error(t, err)
	assert.Nil(t, user1)

	user2, err := InsertUser(goutil.Map{"username": "u", "password": "p"})
	assert.NoError(t, err)
	assert.NotNil(t, user2)

	user3, err := InsertUser(goutil.Map{"username": "u", "password": "p"})
	assert.Error(t, err)
	assert.Nil(t, user3)
}
