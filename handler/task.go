package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/model"
	"github.com/hzxiao/taskmeter/pkg/constvar"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/pkg/httptest"
	"github.com/hzxiao/taskmeter/pkg/timeutil"
	"github.com/lexkong/log"
)

//AddTask POST/api/v1/usr/projects/{project_id}/tasks
func AddTask(c *gin.Context) {
	var task model.Task
	if err := c.Bind(&task); err != nil {
		log.Error("[AddTask} bind data", err)
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	pid := c.Param("project_id")
	//TODO: check pid validity
	_ = pid

	//check task field
	if task.Title == "" {
		SendResponse(c, newArgInvalidError("task's title is empty"), nil)
		return
	}
	//TODO: check tags validity
	//
	task.Uid = c.GetString("uid")
	err := model.InsertTask(&task)
	if err != nil {
		log.Error("[AddTask]", err)
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, goutil.Map{
		"task": task,
	})
}

func DoAddTask(token, pid string, task model.Task) (goutil.Map, error) {
	u := fmt.Sprintf("/api/v1/usr/projects/%v/tasks?token=%v", pid, token)
	return checkResultError(httptest.PostJSON(u, &task))
}

//UpdateTask PUT/api/v1/usr/tasks/{task_id}
func UpdateTask(c *gin.Context) {
	var task model.Task
	if err := c.Bind(&task); err != nil {
		log.Error("[AddTask} bind data", err)
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	task.Id = c.Param("task_id")
	task.Uid = c.GetString("uid")

	oldTask, err := model.LoadTask(task.Id, nil)
	if err != nil {
		log.Error("[UpdateTask]", err)
		SendResponse(c, err, nil)
		return
	}
	if task.Uid != oldTask.Uid {
		SendResponse(c, newArgInvalidError("you aren't the task(%v)'s onwer", task.Id), nil)
		return
	}
	oldState, newState := oldTask.State, task.State
	//
	switch newState {
	case 0:
	case constvar.TaskPaused:
		if oldState == constvar.TaskRunning {
			task.Spending = timeutil.Now() - oldTask.LastStart
		}
	case constvar.TaskRunning:
	case constvar.TaskCompleted:
		if oldState == constvar.TaskRunning {
			task.Spending = timeutil.Now() - oldTask.LastStart
		}
	default:
		SendResponse(c, newArgInvalidError("unknown state(%v)", task.State), nil)
		return
	}

	//TODO: check tags
	//

	err = model.UpdateTaskBasic(task)
	if err != nil {
		log.Error("[UpdateTask]", err)
		SendResponse(c, err, nil)
		return
	}

	//TODO: recode op
	//

	SendResponse(c, nil, goutil.Map{"task": task})
}

//LoadTask GET/api/v1/usr/tasks/{task_id}
func LoadTask(c *gin.Context) {
	tid := c.Param("task_id")
	task, err := model.LoadTask(tid, nil)
	if err != nil {
		log.Error("[LoadTask]", err)
		SendResponse(c, err, nil)
		return
	}

	if task.Uid != c.GetString("uid") {
		SendResponse(c, newArgInvalidError("the task(%v) isn't yours", tid), nil)
		return
	}

	SendResponse(c, nil, goutil.Map{"task": task})
}

//ListTasks GET/api/v1/usr/projects/{project_id}/tasks
func ListTasks(c *gin.Context) {
	argMap, err := CheckURLArg(c.Request.URL.Query(), []*Arg{
		{Key: "state", Type: "int", DefaultValue: "0"},
		{Key: "offset", Type: "int", DefaultValue: "0"},
		{Key: "limit", Type: "int", DefaultValue: "20"},
		{Key: "sortBy", Type: "string"},
		{Key: "sortModel", Type: "string"},
	})
	if err != nil {
		log.Error("[SearchTasks]", err)
		SendResponse(c, err, nil)
		return
	}

	pid := c.Param("project_id")
	_ = pid

	task := model.Task{
		Uid:   c.GetString("uid"),
		State: int(argMap.GetInt64("state")),
	}

	tasks, total, err := model.ListTasks(task, nil, formatSort(argMap.GetString("sortBy"), argMap.GetString("sortModel")),
		int(argMap.GetInt64("offset")), int(argMap.GetInt64("offset")))
	if err != nil {
		log.Error("[ListTasks]", err)
		SendResponse(c, err, nil)
	}

	SendResponse(c, nil, goutil.Map{
		"tasks": tasks,
		"queryInfo": goutil.Map{
			"returnCount": len(tasks),
			"total":       total,
		},
		"cond": argMap,
	})
}

//SearchTasks GET/api/v1/usr/search/tasks
func SearchTasks(c *gin.Context) {
	argMap, err := CheckURLArg(c.Request.URL.Query(), []*Arg{
		{Key: "title", Type: "string"},
		{Key: "state", Type: "int", DefaultValue: "0"},
		{Key: "offset", Type: "int", DefaultValue: "0"},
		{Key: "limit", Type: "int", DefaultValue: "20"},
		{Key: "sortBy", Type: "string"},
		{Key: "sortModel", Type: "string"},
	})
	if err != nil {
		log.Error("[SearchTasks]", err)
		SendResponse(c, err, nil)
		return
	}

	task := model.Task{
		Uid:   c.GetString("uid"),
		Title: argMap.GetString("title"),
		State: int(argMap.GetInt64("state")),
	}

	tasks, total, err := model.ListTasks(task, nil, formatSort(argMap.GetString("sortBy"), argMap.GetString("sortModel")),
		int(argMap.GetInt64("offset")), int(argMap.GetInt64("offset")))
	if err != nil {
		log.Error("[SearchTasks]", err)
		SendResponse(c, err, nil)
	}

	SendResponse(c, nil, goutil.Map{
		"tasks": tasks,
		"queryInfo": goutil.Map{
			"returnCount": len(tasks),
			"total":       total,
		},
		"cond": argMap,
	})
}

//DelTask DELETE/api/v1/usr/tasks/{task_id}
func DelTask(c *gin.Context) {
	tid := c.Param("task_id")
	uid := c.GetString("uid")
	err := model.RemoveTags(uid, []string{tid})
	if err != nil {
		log.Error("[DelTask]", err)
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, nil)
}
