package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connect : connect to the mysql server instance
func Connect() sql.DB {
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		panic(err.Error())
	}

	return db
}
