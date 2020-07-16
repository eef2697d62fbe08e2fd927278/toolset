package tag

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql" // this is needed for mysql
	"github.com/youngtrashbag/toolset/src/utils"
)

// Insert : saves tag to db and returns id of entry
func (t *Tag) Insert() int64 {
	db, err := sql.Open("mysql", "toolset_insert:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	// prepare sql insert note statement
	insertNote, err := db.Prepare("INSERT INTO tbl_tag (name, creationDate) VALUES (?, ?)")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer insertNote.Close()

	var time string
	utils.ConvertTime(&t.CreationDate, &time)

	// execute sql insert note statement
	result, err := insertNote.Exec(t.Name, time)
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

	if tagRows.Err() != nil {
		log.Panicln(tagRows.Err())
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

	if tagRows.Err() != nil {
		log.Panicln(tagRows.Err())
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
