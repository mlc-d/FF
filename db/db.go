package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	dsn = os.Getenv("DSN")
)

func GetDB() *sql.DB {
	if db != nil {
		return db
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
