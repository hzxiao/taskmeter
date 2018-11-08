package model

import (
	"github.com/hzxiao/goutil"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type Model interface {
	Insert() error
	List(selector, sort []string, skip, limit int, needCount bool, result interface{}) (int, error)
	Update(info goutil.Map) error
	Load() error
	Remove() error
}

func insert(coll string, docs ...interface{}) error {
	return DB.C(coll).Insert(docs...)
}

func one(coll string, find, selector bson.M, v interface{}) error {
	return DB.C(coll).Find(find).Select(selector).One(v)
}

func update(coll string, finder, updater bson.M) error {
	return DB.C(coll).Update(finder, updater)
}

func updateAll(coll string, finder, updater bson.M) (*mgo.ChangeInfo, error) {
	return DB.C(coll).UpdateAll(finder, updater)
}

func list(coll string, cond bson.M, selector, sort []string, skip, limit int, needCount bool, v interface{}) (int, error) {
	query := DB.C(coll).Find(cond).Sort(sort...).Select(formatSelector(selector))
	var count int
	var err error
	if needCount {
		count, err = query.Count()
		if err != nil {
			return 0, err
		}
	}

	if skip > 0 {
		query = query.Skip(skip)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	return count, query.All(v)
}

func findAndModify(coll string, finder, updater bson.M, upsert, returnNew, remove bool, result interface{}) (*mgo.ChangeInfo, error) {
	info, err := DB.C(coll).Find(finder).Apply(mgo.Change{
		Update:    updater,
		Upsert:    upsert,
		ReturnNew: returnNew,
		Remove:    remove,
	}, &result)
	return info, err
}

func formatSelector(ss []string) bson.M {
	if len(ss) == 0 {
		return nil
	}

	m := bson.M{}
	for _, s := range ss {
		m[s] = 1
	}
	return m
}

func ContactValue(ss ...string) string {
	return strings.Join(ss, "#")
}
