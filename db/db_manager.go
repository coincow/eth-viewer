package db

import (
	"database/sql"
	"log"
)

var mDB *sql.DB = nil

func GetDB() *sql.DB {
	if nil != mDB {
		return mDB
	}

	log.Println("setup db")
	db, err := SetupDb()
	if nil != err {
		return nil
	} else {
		mDB = db
	}
	return mDB
}
