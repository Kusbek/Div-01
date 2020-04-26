package sqllite

import (
	"DIV-01/new_forum/common"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Connection *sql.DB
}

var DB Database

func Initialize() {
	connection := initDB()
	DB.Connection = connection
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "local_database.db")
	common.AbortOnError(err, "Failed to establish connection with SQL database")
	return db
}

func GetDB() *sql.DB {
	return DB.Connection
}
