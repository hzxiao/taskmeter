package model

import (
	"github.com/hzxiao/goutil"
	"gopkg.in/mgo.v2"
)

const (
	CollSeq  = "taskmeter_seq"
	CollUser = "taskmeter_user"
	CollOp   = "taskmeter_op"
	CollTag  = "taskmeter_tag"
	CollTask = "taskmeter_task"
)

var indexMap = map[string][]mgo.Index{
	CollUser: {
		{Name: "uname", Key: []string{"uname"}, Unique: true},
		{Name: "verification_wxOpenId", Key: []string{"verification.wxOpenId"}, Unique: true, Sparse: true},
	},
	CollTask: {
		{Name: "tags", Key: []string{"tags"}},
		{Name: "running_mark", Key: []string{"runningMark"}, Unique: true, Sparse: true},
	},
	CollTag: {
		{Name: "uidName", Key: []string{"uidName"}, Unique: true, Sparse: true},
	},
}

type Seq struct {
	Id    string `bson:"_id" json:"id"`
	Value uint64 `bson:"value" json:"id"`
}

type User struct {
	Id           string     `bson:"_id" json:"id"`
	UName        []string   `bson:"uname" json:"uname"`
	Password     string     `bson:"password" json:"-"`
	Basic        goutil.Map `bson:"basic" json:"basic"`
	SignUp       goutil.Map `bson:"signUp" json:"-"`
	SignIn       goutil.Map `bson:"signIn" json:"-"`
	Verification goutil.Map `bson:"verification" json:"-"`
	Role         goutil.Map `bson:"role" json:"role"`
	Status       int        `bson:"status" json:"-"`
	Create       int64      `bson:"create" json:"create"`
	Last         int64      `bson:"last" json:"last"`
}

type OpRecord struct {
	Id   string     `bson:"_id" json:"id"`
	Uid  string     `bson:"uid" json:"uid"`
	Type string     `bson:"type" json:"type"`
	Attr goutil.Map `bson:"attr" json:"attr"`
	Time int64      `bson:"time" json:"time"`
}

type Task struct {
	Id          string    `bson:"_id" json:"id"`
	Uid         string    `bson:"uid" json:"uid"`
	Pid         string    `bson:"pid" json:"pid"` //project id
	Title       string    `bson:"title" json:"title"`
	Desc        string    `bson:"desc" json:"desc"`
	Tags        []string  `bson:"tags" json:"tags"`
	Attr        *TaskAttr `bson:"attr" json:"attr"` //tag id list
	Spending    int64     `bson:"spending" json:"spending"`
	LastStart   int64     `bson:"lastStart" json:"lastStart"`
	State       int       `bson:"state" json:"state"`
	Status      int       `bson:"status" json:"-"`
	Create      int64     `bson:"create" json:"create"`
	Last        int64     `bson:"last" json:"last"`
	RunningMark string    `bson:"runningMark" json:"-"` //running mark, unique for the user
}

type TaskAttr struct {
	Duration int64 `bson:"duration" json:"duration"`
	Deadline int64 `bson:"deadline" json:"deadline"`
}

type Tag struct {
	Id      string `bson:"_id" json:"id"`
	Uid     string `bson:"uid" json:"uid"`
	Name    string `bson:"name" json:"name"`
	Status  int    `bson:"status" json:"-"`
	Create  int64  `bson:"create" json:"create"`
	Last    int64  `bson:"last" json:"last"`
	UidName string `bson:"uidName" json:"-"` //unique mask
}

type Project struct {
	Id      string `bson:"_id" json:"id"`
	Uid     string `bson:"uid" json:"uid"`
	Name    string `bson:"name" json:"name"`
	Status  int    `bson:"status" json:"-"`
	Create  int64  `bson:"create" json:"create"`
	Last    int64  `bson:"last" json:"last"`
	UidName string `bson:"uidName" json:"-"` //unique mask
}
