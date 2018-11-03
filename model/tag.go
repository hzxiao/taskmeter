package model

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/hzxiao/taskmeter/pkg/constvar"
	"github.com/hzxiao/taskmeter/pkg/timeutil"
)

func InsertTag(tag *Tag) error {
	if tag == nil {
		return newArgInvalidError("tag is nil")
	}
	tag.Id = "TG"+bson.NewObjectId().Hex()
	tag.Status = constvar.Normal
	tag.Create = timeutil.Now()
	tag.Last = tag.Create
	tag.UidName = ContactValue(tag.Uid, tag.UidName)
	return nil
}

func UpdateTag(tag Tag) error {
	return nil
}

func ListTags(task Task, selector, sort []string, skip, limit int) ([]*Tag, int, error) {
	return nil, 0, nil
}

func LoadTag(id string, selector []string) (*Tag, error) {
	return nil, nil
}

func RemoveTags(uid string, ids []string) error {
	return nil
}

func CheckTagsExist(tags []string) error {
	if len(tags) == 0 {
		return nil
	}

	return nil
}