package model

import (
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/pkg/constvar"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/pkg/timeutil"
	"github.com/lexkong/log"
	"gopkg.in/mgo.v2/bson"
)

func (p *Project) Insert() error {
	if p == nil {
		return newArgInvalidError("project is nil")
	}

	p.Id = "PJ" + bson.NewObjectId().Hex()
	p.Status = constvar.Normal
	p.Create = timeutil.Now()
	p.Last = p.Create
	p.UidName = ContactValue(p.Uid, p.Name)
	err := insert(CollProject, p)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("insert project by data(%v)", goutil.Struct2Json(p))
		log.Error("[Project.Insert]", err)
		return err
	}
	return nil
}

func (p *Project) Load() error {
	var pj Project
	err := one(CollProject, bson.M{"_id": p.Id, "status": constvar.Normal}, nil, &pj)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("load project by id(%v)", p.Id)
		log.Error("[Project.Load]", err)
		return err
	}
	p = &pj
	return nil
}

func (p *Project) List(selector, sort []string, skip, limit int, needCount bool, result interface{}) (int, error) {
	if p == nil {
		return 0, newArgInvalidError("project is nil")
	}
	finder := bson.M{}
	if p.Id != "" {
		finder["_id"] = p.Id
	}
	if p.Uid != "" {
		finder["uid"] = p.Uid
	}
	if p.Status != 0 {
		finder["status"] = p.Status
	}

	total, err := list(CollProject, finder, selector, sort, skip, limit, true, &result)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("list projects by finder(%v)", goutil.Struct2Map(finder))
		log.Error("[Project.List]", err)
		return 0, err
	}
	return total, nil
}

func (p *Project) Update(info goutil.Map) error {
	if p == nil {
		return newArgInvalidError("project is nil")
	}

	finder := bson.M{
		"_id": p.Id,
		"uid": p.Uid,
	}
	set := bson.M{}
	if p.Name != "" {
		set["name"] = p.Name
		set["uidName"] = ContactValue(p.Uid, p.UidName)
	}

	if p.Status != 0 {
		set["status"] = p.Status
	}

	if len(set) > 0 {
		set["last"] = timeutil.Now
	}

	updater := bson.M{}
	if len(set) > 0 {
		updater["$set"] = set
	}

	if len(updater) == 0 {
		return nil
	}

	_, err := findAndModify(CollProject, finder, updater, false, true, false, p)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("update project by finder(%v), updater(%v)", goutil.Struct2Json(finder), goutil.Struct2Json(updater))
		log.Error("[Project.Update]", err)
		return err
	}

	log.Infof("[Project.Update] update tag by finder(%v), updater(%v)", goutil.Struct2Json(finder), goutil.Struct2Json(updater))

	return nil
}

func (p *Project) Remove() error {
	p.Status = constvar.Deleted
	return p.Update(nil)
}
