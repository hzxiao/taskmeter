package model

import "gopkg.in/mgo.v2/bson"

func insert(coll string, docs ...interface{}) error {
	return DB.C(coll).Insert(docs...)
}

func one(coll string, find, selector bson.M, v interface{}) error {
	return DB.C(coll).Find(find).Select(selector).One(&v)
}