package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Connect : connect to the mysql server instance
func Connect() sql.DB {
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		panic(err.Error())
	}

	return *db
}

// ConvertTime : convert the time from go to a string,
//					so it complies with mysql standard for DATETIME.
//					format used is "YYYY-MM-DD hh:mm:ss"
func ConvertTime(tm time.Time) string {
	var t string

	// setting date
	t = string(tm.Year()) + "-"
	if int(tm.Month()) <= 9 {
		t += "0"
	}
	t += string(int(tm.Month())) + "-"
	if int(tm.Day()) <= 9 {
		t += "0"
	}
	t += string(tm.Day()) + " "

	//setting time
	if int(tm.Hour()) <= 9 {
		t += "0"
	}
	t += string(tm.Hour()) + ":"
	if int(tm.Minute()) <= 9 {
		t += "0"
	}
	t += string(tm.Minute()) + ":"
	if int(tm.Second()) <= 9 {
		t += "0"
	}
	t += string(tm.Second())

	return t
}
