package model

import (
	"github.com/hzxiao/taskmeter/config"
	"testing"
)

func init() {
	err := config.Init("../conf/config_test.yaml")
	if err != nil {
		panic(err)
	}

	err = Init()
	if err != nil {
		panic(err)
	}
}

func TestInit(t *testing.T) {

}

func removeAll()  {
	DB.C(CollSeq).RemoveAll(nil)
	DB.C(CollUser).RemoveAll(nil)
	DB.C(CollTask).RemoveAll(nil)
}