package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "P@ssw0rd"
	dbName := "go"
	db, err = sql.Open(dbDriver, dbUser+": "+dbPass+"@/"+dbName)
	return
}
