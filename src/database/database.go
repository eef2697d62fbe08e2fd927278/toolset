package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// InsertConnect : connect to the mysql server with insert privileges
func InsertConnect() sql.DB {
	db, err := sql.Open("mysql", "toolset_insert:password@/toolset")
	if err != nil {
		panic(err.Error())
	}

	return *db
}

// SelectConnect : connect to mysql server with select privileges
func SelectConnect() sql.DB {
	db, err := sql.Open("mysql", "toolset_select:password@/toolset_db")
	if err != nil {
		panic(err.Error())
	}

	return *db
}

// UpdateConnect : connect to mysql server with update privileges
func UpdateConnect() sql.DB {
	db, err := sql.Open("mysql", "toolset_update:password@/toolset_db")
	if err != nil {
		panic(err.Error())
	}

	return *db
}

// DeleteConnect : connect to mysql server with delete privileges
func DeleteConnect() sql.DB {
	db, err := sql.Open("mysql", "toolset_delete:password@/toolset_db")
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
