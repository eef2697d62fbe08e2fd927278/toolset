package note

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // this is needed for mysql
	"github.com/youngtrashbag/toolset/src/database"
)

// Note : Struct used for writing note
type Note struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreationDate time.Time `json:"datetime"`
	AuthorID     int64     `json:"author_id"`
}

// NewNote : a constructor for the Note struct
func NewNote(t string, c string) Note {
	var n Note

	n.Title = t
	n.Content = c
	n.CreationDate = time.Now()

	// TODO: set authorId

	return n
}

// Insert : saves a user in the database
func (n *Note) Insert() int64 {

	// connection to database
	db, err := sql.Open("mysql", "toolset_insert:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	// prepare sql insert note statement
	insertNote, err := db.Prepare("INSERT INTO tbl_note (title, content, time) VALUES (?, ?, ?)")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer insertNote.Close()

	var time string
	database.ConvertTime(&n.CreationDate, &time)
	// execute sql insert note statement
	result, err := insertNote.Exec(n.Title, n.Content, time)
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

// GetByID : returns the selected note from the database as an object
func GetByID(id int64) Note {
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
		err := tagRows.Scan(&n.ID, &n.Title, &n.Content, &timeStr, &n.AuthorID)
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	err = tagRows.Err()
	if err != nil {
		log.Panicln(err.Error())
	}

	if n.ID == 0 && n.Content == "" {
		// when there is no entry found, return id = -1
		n.ID = -1
	}

	database.ConvertTime(&n.CreationDate, &timeStr)
	return n
}

// LinkTag : this links the noteID and the tagId together via the linktable
func LinkTag(nID, tID int64) {
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

	_, err = linkTag.Exec(nID, tID)
	if err != nil {
		log.Panicln(err.Error())
	}

}
