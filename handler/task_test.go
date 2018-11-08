package handler

import (
	"github.com/hzxiao/goutil/assert"
	"github.com/hzxiao/taskmeter/model"
	"testing"
)

func TestAddTask(t *testing.T) {
	removeAll()

	token, err := SignUpAndLogin("Bob")
	assert.NoError(t, err)

	res, err := DoAddTask(token, "cc", model.Task{Title: "Study Go"})
	assert.NoError(t, err)

	_ = res
}
