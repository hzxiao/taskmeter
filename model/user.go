package model

import (
	"fmt"
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/pkg/constvar"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/util"
	"github.com/lexkong/log"
	"gopkg.in/mgo.v2/bson"
)

func InsertUser(user *User) (err error) {
	if user == nil {
		return newArgInvalidError("insert user but user is nil")
	}
	user.Status = constvar.UserNormal
	user.Create = util.Now()
	user.Last = user.Create
	user.Id, err = newUid()
	if err != nil {
		log.Errorf(err, "[InsertUser] arg data(%v)", goutil.Struct2Json(user))
		return err
	}

	err = insert(CollUser, user)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("insert user by data(%v)", goutil.Struct2Json(user))
		log.Errorf(err, "[InsertUser] arg by data(%v)", goutil.Struct2Json(user))
		return err
	}

	return nil
}

func newUid() (string, error) {
	next, err := NextSeq(constvar.UserSeq, 1)
	if err != nil {
		return "", errno.New(errno.ErrDatabase, err).Add("get next seq when new a uid")
	}
	return fmt.Sprintf("u%v", next), nil
}

func FindUser(u *User) (*User, error) {
	if u == nil {
		return nil, newArgInvalidError("find user but u is nil")
	}

	finder := bson.M{}
	if u.Id == "" {
		finder["_id"] = u.Id
	}
	if len(u.UName) > 0 {
		finder["uname"] = bson.M{"$in": u.UName}
	}
	for k, v := range u.Verification {
		finder["verification."+k] = v
	}

	var user User
	err := one(CollUser, finder, nil, &user)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("find user by finder(%v)", goutil.Struct2Json(finder))
		return nil, err
	}

	return &user, nil
}
