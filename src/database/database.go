package database

import (
	//"database/sql"
	"log"
	"strconv"
	"time"
	// _ "github.com/go-sql-driver/mysql"
)

// Response : a struct for json responses
type Response struct {
	Message string `json:"message"`
}

// NewResponse : returns a struct with a message you set as parameter
func NewResponse(m string) Response {
	var r Response
	r.Message = m

	return r
}

// TODO: this file is generally a mess, decide what to do with it and rename it,
//			becuase without the db funcs there is nothing really making this about datbases

// ConvertTime : convert the time from go to a string,
// so it complies with mysql standard for DATETIME.
// format used is "YYYY-MM-DD hh:mm:ss"
func ConvertTime(t *time.Time, s *string) {

	if t.IsZero() {
		var tm time.Time

		var st = *s

		year, err := strconv.Atoi(st[0:4])
		if err != nil {
			log.Panicln(err.Error())
		}
		month, err := strconv.Atoi(st[5:7])
		if err != nil {
			log.Panicln(err.Error())
		}
		day, err := strconv.Atoi(st[8:10])
		if err != nil {
			log.Panicln(err.Error())
		}

		hour, err := strconv.Atoi(st[11:13])
		if err != nil {
			log.Panicln(err.Error())
		}
		minute, err := strconv.Atoi(st[14:16])
		if err != nil {
			log.Panicln(err.Error())
		}
		second, err := strconv.Atoi(st[17:19])
		if err != nil {
			log.Panicln(err.Error())
		}

		tm = time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
		*t = tm
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

		*s = tm
	}
}
