package tag

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/youngtrashbag/toolset/src/database"
)

// Tag : struct for tags, which are used to organise notes
type Tag struct {
	id   int64
	Name string
}

// NewTag : returns a tag object
func NewTag(n string) Tag {
	var t Tag
	t.Name = n

	return t
}

// GetTagByID : returns a tag object from the database
func GetTagByID(id int64) Tag {
	db := database.SelectConnect()
	defer db.Close()

	tagRows, err := db.Query("SELECT id, name FROM tbl.Tag WHERE id = ?;", id)
	if err != nil {
		panic(err.Error())
	}
	defer tagRows.Close()

	var t Tag
	for tagRows.Next() {
		err := tagRows.Scan(&t.id, &t.Name)
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
	db := database.SelectConnect()
	defer db.Close()

	tagRows, err := db.Query("SELECT id, name FROM tbl.Tag WHERE name = ?;", name)
	if err != nil {
		panic(err.Error())
	}
	defer tagRows.Close()

	var t Tag
	for tagRows.Next() {
		err := tagRows.Scan(&t.id, &t.Name)
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
