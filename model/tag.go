package model

import (
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/taskmeter/pkg/constvar"
	"github.com/hzxiao/taskmeter/pkg/errno"
	"github.com/hzxiao/taskmeter/pkg/timeutil"
	"github.com/lexkong/log"
	"gopkg.in/mgo.v2/bson"
)

func InsertTag(tag *Tag) error {
	if tag == nil {
		return newArgInvalidError("tag is nil")
	}
	tag.Id = "TG" + bson.NewObjectId().Hex()
	tag.Status = constvar.Normal
	tag.Create = timeutil.Now()
	tag.Last = tag.Create
	tag.UidName = ContactValue(tag.Uid, tag.UidName)

	err := insert(CollTag, tag)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("insert tag by data(%v)", goutil.Struct2Json(tag))
		log.Error("[InsertTag]", err)
		return err
	}
	return nil
}

func UpdateTag(tag *Tag) error {
	if tag == nil {
		return newArgInvalidError("tag is nil")
	}

	finder := bson.M{
		"_id": tag.Id,
		"uid": tag.Uid,
	}

	set := bson.M{}
	if tag.Name != "" {
		set["name"] = tag.Name
		set["uidName"] = ContactValue(tag.Uid, tag.UidName)
		set["last"] = timeutil.Now
	}

	updater := bson.M{}
	if len(set) > 0 {
		updater["$set"] = set
	}

	if len(updater) == 0 {
		return nil
	}

	_, err := findAndModify(CollTag, finder, updater, false, true, false, tag)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("update tag by finder(%v), updater(%v)", goutil.Struct2Json(finder), goutil.Struct2Json(updater))
		log.Error("[UpdateTag]", err)
		return err
	}

	return nil
}

func ListTags(tag Tag, selector, sort []string, skip, limit int) ([]*Tag, int, error) {
	finder := bson.M{}
	if tag.Id != "" {
		finder["_id"] = tag.Id
	}
	if tag.Uid != "" {
		finder["uid"] = tag.Uid
	}
	if tag.Status != 0 {
		finder["status"] = tag.Status
	}

	var tags []*Tag
	total, err := list(CollTag, finder, selector, sort, skip, limit, true, &tags)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("list tags by finder(%v)", goutil.Struct2Map(finder))
		log.Error("[ListTags]", err)
		return nil, 0, err
	}
	return tags, total, nil
}

func LoadTag(id string, selector []string) (*Tag, error) {
	var tag Tag
	err := one(CollTag, bson.M{"_id": id, "status": constvar.Normal}, formatSelector(selector), &tag)
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("load tag by id(%v)", id)
		log.Error("[LoadTag]", err)
		return nil, err
	}
	return &tag, nil
}

func RemoveTags(uid string, ids []string) error {
	finder := bson.M{
		"uid": uid,
		"_id": bson.M{"$in": ids},
	}

	set := bson.M{}
	set["status"] = constvar.Deleted
	set["last"] = timeutil.Now()

	_, err := updateAll(CollTag, finder, bson.M{"$set": set})
	if err != nil {
		err = errno.New(errno.ErrDatabase, err).Addf("remove tags by finder(%v)", goutil.Struct2Json(finder))
		log.Error("[RemoveTags]", err)
		return err
	}
	return nil
}

func CheckTagsExist(tags []string) error {
	if len(tags) == 0 {
		return nil
	}

	return nil
}
