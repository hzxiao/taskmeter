package model

import (
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/pkg/constvar"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/pkg/timeutil"
	"github.com/lexkong/log"
	"gopkg.in/mgo.v2/bson"
)

func InsertTask(task Task) error {
	task.Id = "TK" + bson.NewObjectId().Hex()
	task.Spending = 0
	task.LastStart = 0
	task.Status = constvar.Normal
	task.State = constvar.TaskPaused
	task.Create = timeutil.Now()
	task.Last = task.Create

	err := insert(CollTask, &task)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("insert task by data(%gv)", goutil.Struct2Json(&task))
		log.Errorf(err, "[InsertTask] %v", err)
		return err
	}
	return nil
}

func UpdateTaskBasic(task Task) error {
	finder := bson.M{
		"_id": task.Id,
		"uid": task.Uid,
	}

	set, unset, inc := bson.M{}, bson.M{}, bson.M{}
	if task.Pid != "" {
		set["pid"] = task.Pid
	}
	if task.Title != "" {
		set["title"] = task.Title
	}
	if task.Desc != "" {
		set["desc"] = task.Desc
	}
	if task.Tags != nil {
		set["tags"] = task.Tags
	}
	if task.Attr != nil {
		set["attr"] = task.Attr
	}
	if task.Status > 0 {
		set["status"] = task.Status
	}
	if task.State == constvar.TaskPaused {
		set["state"] = task.State
		//清除运行标志
		unset["runningMark"] = true
		//只暂停正在执行的任务
		finder["runningMark"] = task.Uid
		if task.Spending > 0 { //增加花费的时长
			inc["spending"] = task.Spending
		}
	} else if task.State == constvar.TaskRunning {
		set["state"] = task.State
		set["lastStart"] = timeutil.Now()
		set["runningMark"] = task.Uid
		//
		finder["runningMark"] = bson.M{"$exist": false}
	}

	if len(set) == 0 {
		return newArgInvalidError("nothing to update")
	}
	set["last"] = timeutil.Now()

	updater := bson.M{"$set": set}
	if len(unset) > 0 {
		updater["$unset"] = unset
	}
	if len(inc) > 0 {
		updater["$inc"] = inc
	}

	err := update(CollTask, finder, updater)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err)
		log.Errorf(err, "[UpdateTaskBasic] update task by finder(%v), updater(%v)", goutil.Struct2Json(finder), goutil.Struct2Json(updater))
		return err
	}
	return nil
}

func UpdateTaskExtra(task Task) error {
	set, unset, inc := bson.M{}, bson.M{}, bson.M{}
	if task.Spending > 0 {
		inc["spending"] = task.Spending
	}
	if task.State == constvar.TaskPaused || task.State == constvar.TaskOverdue {
		set["state"] = task.State
		unset["runningMark"] = true
	}

	finder := bson.M{"_id": task.Id}
	updater := bson.M{}
	if len(set) > 0 {
		updater["$set"] = set
	}
	if len(unset) > 0 {
		updater["$unset"] = unset
	}
	if len(inc) > 0 {
		updater["$inc"] = inc
	}

	return update(CollTask, finder, updater)
}

func ListTasks(task Task, selector, sort []string, skip, limit int) ([]*Task, int, error) {
	finder := bson.M{}
	if task.Id != "" {
		finder["_id"] = task.Id
	}
	if task.Uid != "" {
		finder["uid"] = task.Uid
	}
	if task.Pid != "" {
		finder["pid"] = task.Pid
	}
	if len(task.Tags) > 0 {

	}
	return nil, 0, nil
}

func LoadTask(id string, selector []string) (*Task, error) {
	return nil, nil
}
