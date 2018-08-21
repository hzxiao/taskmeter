package model

import (
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/util"
	"github.com/lexkong/log"
	"gopkg.in/mgo.v2/bson"
)

func InsertOpRecord(r *OpRecord) error {
	if r == nil {
		return newArgInvalidError("insert op record: r is nil")
	}
	if r.Uid == "" {
		return newArgInvalidError("insert op record: uid is empty")
	}
	if r.Type == "" {
		return newArgInvalidError("insert op record: type is empty")
	}

	r.Id = "OR" + bson.NewObjectId().Hex()
	r.Time = util.Now()
	err := insert(CollOp, r)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("insert op by r(%v)", goutil.Struct2Json(r))
		log.Errorf(err, "[InserOpRecord] arg by r(%v)", goutil.Struct2Json(r))
		return err
	}

	return nil
}
