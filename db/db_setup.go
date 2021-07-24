package db

import (
	"database/sql"
	_ "database/sql"
)
import _ "github.com/go-sql-driver/mysql"

// Setup the database connection
func DBSetup() (*sql.DB, error) {
	dbUser := "root"
	dbPasswd := "Ashu_000"
	dbHost := "47.98.54.147:3306"
	dbName := "mysql"
	dbConnection := "tcp"
	connectionString := dbUser + ":" + dbPasswd + "@" + dbConnection + "(" + dbHost + ")" + "/" + dbName + "?charset=utf8"

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		print("sql open:" + err.Error())
		return nil, err
	}

	// Ping the database once since Open() doesn't open a connection
	err = db.Ping()
	if err != nil {
		print("do.ping:" + err.Error())
		return nil, err
	}

	return db, nil
}
