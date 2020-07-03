package tag

import (
	"database/sql"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql" // this is needed for mysql
)

// Tag : struct for tags, which are used to organise notes
type Tag struct {
	ID           int64
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creation_date"`
}

// NewTag : returns a tag object
func NewTag(n string) Tag {
	var t Tag
	t.Name = n

	return t
}

// Insert : saves tag to db and returns id of entry
func (t *Tag) Insert() int64 {
	db, err := sql.Open("mysql", "toolset_insert:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	// prepare sql insert note statement
	insertNote, err := db.Prepare("INSERT INTO tbl_tag (name) VALUES (?)")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer insertNote.Close()

	// execute sql insert note statement
	result, err := insertNote.Exec(t.Name)
	if err != nil {
		log.Panicln(err.Error())
	}

	tagID, err := result.LastInsertId()
	if err != nil {
		log.Panicln(err.Error())
	}

	return tagID
}

// GetByID : returns a tag object from the database
func GetByID(id int64) Tag {
	db, err := sql.Open("mysql", "toolset_select:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	tagRows, err := db.Query("SELECT id, name FROM tbl.Tag WHERE id = ?", id)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer tagRows.Close()

	var t Tag
	for tagRows.Next() {
		err := tagRows.Scan(&t.ID, &t.Name)
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	err = tagRows.Err()
	if err != nil {
		log.Panicln(err.Error())
	}

	return t
}

// GetByName : return a tag searched for by its name
func GetByName(name string) Tag {
	db, err := sql.Open("mysql", "toolset_select:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	name = strings.ToLower(name)
	tagRows, err := db.Query("SELECT id, name FROM tbl.Tag WHERE name = ?", name)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer tagRows.Close()

	var t Tag
	for tagRows.Next() {
		err := tagRows.Scan(&t.ID, &t.Name)
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	err = tagRows.Err()
	if err != nil {
		log.Panicln(err.Error())
	}

	return t
}

// LinkNote : this links the noteID and the tagId together via the linktable
func LinkNote(nID, tID int64) {
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
