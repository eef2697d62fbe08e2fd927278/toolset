package database

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// TODO: maybe remove these functions, becuase really it just adds another layer of complexity

// InsertConnect : connect to the mysql server with insert privileges
// note: defer closing the db object after assigning it
func InsertConnect() sql.DB {
	db, err := sql.Open("mysql", "toolset_insert:password@/toolset")
	if err != nil {
		panic(err.Error())
	}

	return *db
}

// SelectConnect : connect to mysql server with select privileges
// note: defer closing the db object after assigning it
func SelectConnect() sql.DB {
	db, err := sql.Open("mysql", "toolset_select:password@/toolset_db")
	if err != nil {
		panic(err.Error())
	}

	return *db
}

// UpdateConnect : connect to mysql server with update privileges
// note: defer closing the db object after assigning it
func UpdateConnect() sql.DB {
	db, err := sql.Open("mysql", "toolset_update:password@/toolset_db")
	if err != nil {
		panic(err.Error())
	}

	return *db
}

// DeleteConnect : connect to mysql server with delete privileges
// note: defer closing the db object after assigning it
func DeleteConnect() sql.DB {
	db, err := sql.Open("mysql", "toolset_delete:password@/toolset_db")
	if err != nil {
		panic(err.Error())
	}

	return *db
}

// ConvertTime : convert the time from go to a string,
// so it complies with mysql standard for DATETIME.
// format used is "YYYY-MM-DD hh:mm:ss"
func ConvertTime(t *time.Time, s *string) {

	if t.IsZero() {
		var tm time.Time

		var st = *s

		year, err := strconv.Atoi(st[0:3])
		if err != nil {
			panic(err.Error())
		}
		month, err := strconv.Atoi(st[5:6])
		if err != nil {
			panic(err.Error())
		}
		day, err := strconv.Atoi(st[8:9])
		if err != nil {
			panic(err.Error())
		}

		hour, err := strconv.Atoi(st[11:12])
		if err != nil {
			panic(err.Error())
		}
		minute, err := strconv.Atoi(st[14:15])
		if err != nil {
			panic(err.Error())
		}
		second, err := strconv.Atoi(st[17:18])
		if err != nil {
			panic(err.Error())
		}

		tm = time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
		t = &tm

	} else if len(*s) <= 0 {
		var tm string

		// setting date
		tm = string(t.Year()) + "-"
		if int(t.Month()) <= 9 {
			tm += "0"
		}
		tm += string(int(t.Month())) + "-"
		if int(t.Day()) <= 9 {
			tm += "0"
		}
		tm += string(t.Day()) + " "

		//setting time
		if int(t.Hour()) <= 9 {
			tm += "0"
		}
		tm += string(t.Hour()) + ":"
		if int(t.Minute()) <= 9 {
			tm += "0"
		}
		tm += string(t.Minute()) + ":"
		if int(t.Second()) <= 9 {
			tm += "0"
		}
		tm += string(t.Second())

		s = &tm
	}
}
