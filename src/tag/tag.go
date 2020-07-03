package tag

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Tag : struct for tags, which are used to organise notes
type Tag struct {
	id   int64
	name string
}

// NewTag : returns a tag object
func NewTag(n string) Tag {
	var t Tag
	t.name = n

	return t
}

// Insert : saves tag to db
func (t *Tag) Insert() int64 {
	db, err := sql.Open("mysql", "toolset_insert:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	// prepare sql insert note statement
	insertNote, err := db.Prepare("INSERT INTO tbl_tag (name) VALUES (?)")
	if err != nil {
		panic(err.Error())
	}
	defer insertNote.Close()

	// execute sql insert note statement
	result, err := insertNote.Exec(t.name)
	if err != nil {
		panic(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return id
}

// GetTagByID : returns a tag object from the database
func GetTagByID(id int64) Tag {
	db, err := sql.Open("mysql", "toolset_select:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	tagRows, err := db.Query("SELECT id, name FROM tbl.Tag WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer tagRows.Close()

	var t Tag
	for tagRows.Next() {
		err := tagRows.Scan(&t.id, &t.name)
		if err != nil {
			panic(err.Error())
		}
	}

	err = tagRows.Err()
	if err != nil {
		panic(err.Error())
	}

	return t
}

// GetTagByName : return a tag searched for by its name
func GetTagByName(name string) Tag {
	db, err := sql.Open("mysql", "toolset_select:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	name = strings.ToLower(name)
	tagRows, err := db.Query("SELECT id, name FROM tbl.Tag WHERE name = ?", name)
	if err != nil {
		panic(err.Error())
	}
	defer tagRows.Close()

	var t Tag
	for tagRows.Next() {
		err := tagRows.Scan(&t.id, &t.name)
		if err != nil {
			panic(err.Error())
		}
	}

	err = tagRows.Err()
	if err != nil {
		panic(err.Error())
	}

	return t
}
