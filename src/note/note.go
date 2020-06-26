package note

import (
	"database/sql"
	"time"
)

// Note : Struct used for writing note
type Note struct {
	Title   string
	Content string
	time    time.Time
	Tags    []string
}

// NewNote : a constructor for the Note struct
func NewNote(t string, c string, tg []string) Note {
	var n Note

	n.Title = t
	n.Content = c
	n.SetTime(time.Now())
	n.Tags = tg

	return n
}

// SetTime : sets time of note
func (n *Note) SetTime(t time.Time) {
	n.time = t
}

// _convertTime : convert the time from go to a string,
//					so it complies with mysql standard for DATETIME.
//					format used is "YYYY-MM-DD hh:mm:ss"
func (n Note) _convertTime() string {
	var t string

	// setting date
	t = string(n.time.Year()) + "-"
	if int(n.time.Month()) <= 9 {
		t += "0"
	}
	t += string(int(n.time.Month())) + "-"
	if int(n.time.Day()) <= 9 {
		t += "0"
	}
	t += string(n.time.Day()) + " "

	//setting time
	if int(n.time.Hour()) <= 9 {
		t += "0"
	}
	t += string(n.time.Hour()) + ":"
	if int(n.time.Minute()) <= 9 {
		t += "0"
	}
	t += string(n.time.Minute()) + ":"
	if int(n.time.Second()) <= 9 {
		t += "0"
	}
	t += string(n.time.Second())

	return "2020-06-26 15:30:00"
}

// Insert : saves a user in the database
func (n Note) Insert() {

	// connection to database TODO: move this to seperate file (database.go) so everything is organized
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// prepare sql insert note statement
	insertUser, err := db.Prepare("INSERT INTO tbl_note (title, content, time) VALUES (?, ?, ?")
	if err != nil {
		panic(err.Error())
	}
	defer insertUser.Close()

	// execute sql insert note statement
	noteInsert, err := insertUser.Exec(n.Title, n.Content, n._convertTime())
	if err != nil {
		panic(err.Error())
	}

	//get the id of inserted note
	id, _ := noteInsert.LastInsertId()

	// prepare sql insert statement into link table (for tags)
	insertTags, err := db.Prepare("INSERT INTO lktbl_tags (note_id, tag) VALUES (?, ?)")
	defer insertTags.Close()

	// insert a row for each tag
	for i := 0; i < len(n.Tags); i++ {
		_, err = insertTags.Exec(id, n.Tags[i])
		if err != nil {
			panic(err.Error())
		}
	}
}
