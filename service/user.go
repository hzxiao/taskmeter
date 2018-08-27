package service

import (
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/model"
	"github.com/hzxiao/taskmeter/pkg/constvar"
	"github.com/hzxiao/taskmeter/util"
	"github.com/lexkong/log"
)

func SignUp(data goutil.Map) (user *model.User, err error) {
	err = checkSignUpData(data)
	if err != nil {
		log.Errorf(err, "[SignUp] arg data(%v)", goutil.Struct2Json(data))
		return nil, err
	}

	user = &model.User{SignUp: data}
	switch data.GetStringP("method") {
	case constvar.WeiXinRegister:
		user.Verification = goutil.Map{
			"wxOpenId": data.GetString("wxOpenId"),
		}
	case constvar.UsernameRegister:
		user.UName = []string{data.GetString("username")}
		user.Password = util.Sha256([]byte(data.GetString("password")))
	}

	err = model.InsertUser(user)
	if err != nil {
		log.Errorf(err, "[SignUp] arg by data(%v)", goutil.Struct2Json(data))
		return nil, err
	}

	return user, nil
}

func checkSignUpData(data goutil.Map) error {
	if data == nil {
		return newArgInvalidError("check sign up data: data is nil")
	}
	switch data.GetString("method") {
	case constvar.WeiXinRegister:
		if data.GetString("wxOpenId") == "" {
			return newArgInvalidError("check sign up data: wxOpenId is empty")
		}
	case constvar.UsernameRegister:
		if data.GetString("username") == "" {
			return newArgInvalidError("check sign up data: username is empty")
		}
		if data.GetString("password") == "" {
			return newArgInvalidError("check sign up data: password is empty")
		}
	default:
		return newArgInvalidError("check sign up data: unknown method(%v)", data.GetString("method"))
	}

	return nil
}

func Login(username, password string) (goutil.Map, error) {
	return nil, nil
}

func WXLogin(code string) (goutil.Map, error) {
	return nil, nil
}