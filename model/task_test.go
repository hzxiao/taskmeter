package model

import (
	"github.com/hzxiao/goutil/assert"
	"github.com/hzxiao/taskmeter/pkg/constvar"
	"testing"
)

func TestInsertTask(t *testing.T) {
	removeAll()

	t1 := Task{}
	err := InsertTask(&t1)
	assert.NoError(t, err)

	//test error
	err = InsertTask(nil)
	assert.Error(t, err)
}

func TestUpdateTaskBasic(t *testing.T) {
	removeAll()

	uid := "uu"
	var err error

	task := Task{
		Uid:   uid,
		Title: "first task",
		Desc:  "description for the task",
		Attr: &TaskAttr{
			Deadline: 0,
			Duration: 0,
		},
	}
	err = InsertTask(&task)
	assert.NoError(t, err)

	//1. update title, desc, tags, attr, pid
	upTask1 := Task{
		Id:    task.Id,
		Uid:   uid,
		Title: "new title for task",
		Desc:  "new desc for task",
		Tags: []string{
			"TG1",
			"TG2",
		},
		Attr: &TaskAttr{
			Duration: 12,
			Deadline: 100,
		},
		Pid: "PJ1",
	}
	err = UpdateTaskBasic(upTask1)
	assert.NoError(t, err)

	realTask1, err := LoadTask(task.Id, nil)
	assert.NoError(t, err)
	assert.Equal(t, upTask1.Title, realTask1.Title)
	assert.Equal(t, upTask1.Desc, realTask1.Desc)
	assert.Equal(t, upTask1.Tags, realTask1.Tags)
	assert.Equal(t, upTask1.Attr, realTask1.Attr)
	assert.Equal(t, upTask1.Pid, realTask1.Pid)

	//2. start the task
	upTask2 := Task{
		Id:    task.Id,
		Uid:   uid,
		State: constvar.TaskRunning,
	}
	err = UpdateTaskBasic(upTask2)
	assert.NoError(t, err)

	realTask2, err := LoadTask(task.Id, nil)
	assert.NoError(t, err)
	assert.Equal(t, uid, realTask2.RunningMark)
	assert.Equal(t, constvar.TaskRunning, realTask2.State)
	assert.NotEqual(t, 0, realTask2.LastStart)

	//dup start error
	err = UpdateTaskBasic(upTask2)
	assert.Error(t, err)

	//3. stop the task
	upTask3 := Task{
		Id:       task.Id,
		Uid:      uid,
		State:    constvar.TaskPaused,
		Spending: 100,
	}
	err = UpdateTaskBasic(upTask3)
	assert.NoError(t, err)

	realTask3, err := LoadTask(task.Id, nil)
	assert.NoError(t, err)

	assert.Equal(t, constvar.TaskPaused, realTask3.State)
	assert.Equal(t, "", realTask3.RunningMark)
	assert.Equal(t, upTask3.Spending, realTask3.Spending)

	//dup stop error
	err = UpdateTaskBasic(upTask3)
	assert.Error(t, err)

	//4. delete the task
	upTask4 := Task{
		Id:     task.Id,
		Uid:    uid,
		Status: constvar.Deleted,
	}
	err = UpdateTaskBasic(upTask4)
	assert.NoError(t, err)

	_, err = LoadTask(task.Id, nil)
	assert.Error(t, err)

}

func TestUpdateTaskExtra(t *testing.T) {
	removeAll()
}

func TestListTasks(t *testing.T) {
	removeAll()

}

func TestLoadTask(t *testing.T) {
	removeAll()

}
