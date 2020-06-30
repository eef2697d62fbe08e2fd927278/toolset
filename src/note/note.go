package note

import (
	"time"

	"github.com/youngtrashbag/toolset/src/database"
	"github.com/youngtrashbag/toolset/src/tag"
)

// Note : Struct used for writing note
type Note struct {
	id      int64
	Title   string
	Content string
	time    time.Time
	Tags    []tag.Tag
}

// NewNote : a constructor for the Note struct
func NewNote(t string, c string, tg []tag.Tag) Note {
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

// GetTime : returns the time of the note
func (n *Note) GetTime() time.Time {
	return n.time
}

// Insert : saves a user in the database
func (n *Note) Insert() {

	// connection to database
	db := database.InsertConnect()
	defer db.Close()

	// prepare sql insert note statement
	insertUser, err := db.Prepare("INSERT INTO tbl_note (title, content, time) VALUES (?, ?, ?);")
	if err != nil {
		panic(err.Error())
	}
	defer insertUser.Close()

	// execute sql insert note statement
	noteInsert, err := insertUser.Exec(n.Title, n.Content, database.ConvertTime(n.GetTime()))
	if err != nil {
		panic(err.Error())
	}

	//get the id of inserted note
	id, _ := noteInsert.LastInsertId()

	// prepare sql insert statement into link table (for tags)
	insertTags, err := db.Prepare("INSERT INTO lktbl_tags (note_id, tag) VALUES (?, ?);")
	defer insertTags.Close()

	// insert a row for each tag
	for i := 0; i < len(n.Tags); i++ {
		_, err = insertTags.Exec(id, n.Tags[i])
		if err != nil {
			panic(err.Error())
		}
	}
}

// TODO: if its called by using `note.GetNoteByID(1)` will it be a bad redundancy in the name ?

// GetNoteByID : returns the selected note from the database as an object
func GetNoteByID(id int64) Note {
	db := database.SelectConnect()
	defer db.Close()

	return NewNote("", "content1", []tag.Tag{tag.NewTag("tag1"), tag.NewTag("tag2")})
}

// GetNoteByTagID : returns a note based on tag
func GetNoteByTagID(tagID int64) Note {
	return NewNote("", "content1", []tag.Tag{tag.NewTag("tag1"), tag.NewTag("tag2")})
}
