package model

import (
	"github.com/hzxiao/goutil"
	"gopkg.in/mgo.v2"
)

const (
	CollSeq  = "taskmeter_seq"
	CollUser = "taskmeter_user"
	CollOp   = "taskmeter_op"
)

var indexMap = map[string][]mgo.Index{
	CollUser: {
		{Name: "uname", Key: []string{"uname"}, Unique: true},
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
