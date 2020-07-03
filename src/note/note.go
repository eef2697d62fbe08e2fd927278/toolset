package note

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/youngtrashbag/toolset/src/database"
)

// Note : Struct used for writing note
type Note struct {
	id           int64     `json: id`
	title        string    `json: title`
	content      string    `json: content`
	creationDate time.Time `json: datetime`
	authorId     int64     `json: author_id`
}

// NewNote : a constructor for the Note struct
func NewNote(t string, c string) Note {
	var n Note

	n.title = t
	n.content = c
	n.time = time.Now()

	// TODO: set authorId

	return n
}

// Insert : saves a user in the database
func (n *Note) Insert() int64 {

	// connection to database
	db, err := sql.Open("mysql", "toolset_insert:password@/toolset")
	if err != nil {
		log.Panicln(err.Error)
	}
	defer db.Close()

	// prepare sql insert note statement
	insertNote, err := db.Prepare("INSERT INTO tbl_note (title, content, time) VALUES (?, ?, ?)")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer insertNote.Close()

	var time string
	database.ConvertTime(&n.time, &time)
	// execute sql insert note statement
	result, err := insertNote.Exec(n.title, n.content, time)
	if err != nil {
		log.Panicln(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln(err.Error())
	}

	return id

	// TODO: fix the problem with tags
	//get the id of inserted note
	//id, _ := noteInsert.LastInsertId()
	// prepare sql insert statement into link table (for tags)
	//insertTags, err := db.Prepare("INSERT INTO lktbl_tags (note_id, tag) VALUES (?, ?);")
	//defer insertTags.Close()
	// insert a row for each tag
	//for i := 0; i < len(n.Tags); i++ {
	//	_, err = insertTags.Exec(id, n.Tags[i])
	//	if err != nil {
	//		log.Panicln(err.Error())
	//	}
	//}
}

// GetById : returns the selected note from the database as an object
func GetById(id int64) Note {
	db, err := sql.Open("mysql", "toolset_select:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	tagRows, err := db.Query("SELECT id, title, content, time, author FROM tbl_note WHERE id = ?", id)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer tagRows.Close()

	var n Note
	var timeStr string
	for tagRows.Next() {
		err := tagRows.Scan(&n.id, &n.title, &n.content, &timeStr, &n.authorId)
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	err = tagRows.Err()
	if err != nil {
		log.Panicln(err.Error())
	}

	if n.id == 0 && n.content == "" {
		// when there is no entry found, return id = -1
		n.id = -1
	}

	database.ConvertTime(&n.time, &timeStr)
	return n
}

// LinkTag : this links the noteID and the tagId together via the linktable
func LinkTag(nId, tId int64) {
	db, err := sql.Open("mysql", "toolset_insert:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	linkTag, err := db.Prepare("INSERT INTO lktbl_tag (note_id, tag_id) VALUES (?, ?)")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer linkTag.Close()

	_, err := linkTag.Exec(nId, tId)
	if err != nil {
		log.Panicln(err.Error())
	}

}
