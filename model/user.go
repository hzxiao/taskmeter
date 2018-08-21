package model

import (
	"fmt"
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/pkg/constvar"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/util"
	"github.com/lexkong/log"
)

func InsertUser(data goutil.Map) (user *User, err error) {
	err = checkSignUpData(data)
	if err != nil {
		log.Errorf(err, "[InsertUser] arg data(%v)", goutil.Struct2Json(data))
		return nil, err
	}

	user = &User{
		UName:    []string{data.GetString("username")},
		Password: util.Sha256([]byte(data.GetString("password"))),
		SignUp:   data,
		Status:   constvar.UserNormal,
		Create:   util.Now(),
		Last:     util.Now(),
	}
	user.Id, err = newUid()
	if err != nil {
		log.Errorf(err, "[InsertUser] arg data(%v)", goutil.Struct2Json(data))
		return nil, err
	}

	err = insert(CollUser, user)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("insert user by data(%v)", goutil.Struct2Json(user))
		log.Errorf(err, "[InsertUser] arg by data(%v)", goutil.Struct2Json(data))
		return nil, err
	}

	return user, nil
}

func checkSignUpData(data goutil.Map) error {
	if data == nil {
		return newArgInvalidError("check sign up data: data is nil")
	}
	if data.GetString("username") == "" {
		return newArgInvalidError("check sign up data: username is empty")
	}
	if data.GetString("password") == "" {
		return newArgInvalidError("check sign up data: password is empty")
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
