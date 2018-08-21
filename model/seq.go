package model

import (
	"github.com/lexkong/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NextSeq(id string, delta uint64) (uint64, error) {
	var seq Seq
	_, err := DB.C(CollSeq).FindId(id).Apply(mgo.Change{
		Update:    bson.M{"$inc": bson.M{"value": delta}},
		Upsert:    true,
		ReturnNew: true,
	}, &seq)
	if err != nil {
		log.Errorf(err, "[NextSeq] id(%v), delta(%v)", id, delta)
		return 0, err
	}
	return seq.Value, nil
}

