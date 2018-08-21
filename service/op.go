package service

import (
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/model"
	"github.com/hzxiao/taskmeter/pkg/constvar"
	"github.com/lexkong/log"
)

func AddSignInRecord(uid string, attr goutil.Map) error {
	r := &model.OpRecord{
		Type: constvar.SignInOp,
		Uid:  uid,
		Attr: attr,
	}

	err := model.InsertOpRecord(r)
	if err != nil {
		log.Errorf(err, "[AddSignInRecord] add sign in record")
		return err
	}
	return nil
}
