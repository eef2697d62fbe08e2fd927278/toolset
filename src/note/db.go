package note

import (
	"database/sql"
	"log"

	"github.com/youngtrashbag/toolset/src/database"
	"github.com/youngtrashbag/toolset/src/tag"
)

// Insert : saves a user in the database and returns the id of said db entry
func (n *Note) Insert() int64 {

	// connection to database
	db, err := sql.Open("mysql", "toolset_insert:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	// prepare sql insert note statement
	insertNote, err := db.Prepare("INSERT INTO tbl_note (title, content, creationDate) VALUES (?, ?, ?)")
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

	noteID, err := result.LastInsertId()
	if err != nil {
		log.Panicln(err.Error())
	}

	// TODO: actually take the tags from somewhere real
	// the tags as strings
	var tagsS []string

	for _, t := range tagsS {
		tg := tag.GetByName(t)

		tID := tg.ID
		// if the note is not yet in the db it will be inserted
		if tID == -1 {
			tID = tg.Insert()
		}
		tag.LinkNote(noteID, tID)
	}

	return noteID
}

// GetByID : returns the selected note from the database as an object
func GetByID(id int64) Note {
	db, err := sql.Open("mysql", "toolset_select:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	noteRows, err := db.Query("SELECT id, title, content, creationDate, author FROM tbl_note WHERE id = ?", id)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer noteRows.Close()

	var n Note
	var timeStr string
	for noteRows.Next() {
		err := noteRows.Scan(&n.ID, &n.Title, &n.Content, &timeStr, &n.AuthorID)
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	if noteRows.Err() != nil {
		log.Panicln(noteRows.Err())
	}

	if n.ID == 0 && n.Content == "" {
		// when there is no entry found, return id = -1
		n.ID = -1
	}

	database.ConvertTime(&n.CreationDate, &timeStr)
	return n
}
