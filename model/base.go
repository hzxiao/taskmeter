package model

func insert(coll string, docs ...interface{}) error {
	return DB.C(coll).Insert(docs...)
}
