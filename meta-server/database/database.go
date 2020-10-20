package database

import (
	"database/sql"
	"log"

	// Implementation of the database/sql
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:password@/oss_demo")
	// db, err := sql.Open("mysql", "root:password@/oss_demo") 这是临时变量db
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

// GetDatabase method.
func GetDatabase() *sql.DB {
	return db
}
