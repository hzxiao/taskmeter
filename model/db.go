package model

import (
	"errors"
	"github.com/hzxiao/taskmeter/config"
	"github.com/lexkong/log"
	"gopkg.in/mgo.v2"
	"time"
)

var DB *Database

type Database struct {
	Session *mgo.Session
	DB      *mgo.Database
	C       func(name string) *mgo.Collection
}

func Init() error {
	var err error
	DB, err = openDB(config.GetString("db.addr"), config.GetString("db.name"))
	return err
}

func openDB(addr string, dbName string) (*Database, error) {
	db := &Database{}
	var err error
	db.Session, err = mgo.Dial(addr)
	if err != nil {
		return db, err
	}

	db.DB = db.Session.DB(dbName)
	db.C = db.DB.C

	go db.pingLoop()
	return db, nil
}

func (db *Database) pingLoop() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		<-ticker.C
		err := db.ping()
		if err == nil {
			continue
		}
		//handle err
		for {
			err = Init()
			if err != nil {
				log.Errorf(err, "try to dial mongo by url(%v) fail.", config.GetString("db.addr"))
				time.Sleep(5 * time.Second)
				continue
			}
			log.Infof("reconnect to mongo success by url(%v)\n", config.GetString("db.addr"))
			return
		}
	}
}

func (db *Database) ping() (err error) {
	errClosed := errors.New("Closed explicitly")
	defer func() {
		if pe := recover(); pe != nil {
			if db.Session != nil {
				db.Session.Clone()
				err = errClosed
			}
		}
	}()

	err = db.Session.Ping()
	if err == nil {
		return nil
	}
	if err.Error() == "Closed explicitly" || err.Error() == "EOF" {
		db.Session.Clone()
		return errClosed
	}
	return
}
