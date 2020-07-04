package user

import (
	"database/sql"
	"log"

	"github.com/youngtrashbag/toolset/src/database"
)

// Insert : saves a user in the database
func (u *User) Insert() int64 {
	// connection to database
	db, err := sql.Open("mysql", "toolset_insert:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	// prepare sql statement
	insertUser, err := db.Prepare("INSERT INTO tbl_user (email, username, password, creationDate) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer insertUser.Close()

	var time string
	database.ConvertTime(&u.CreationDate, &time)
	// execute sql statement
	result, err := insertUser.Exec(u.Email, u.Username, u.password, time)
	if err != nil {
		log.Panicln(err.Error())
	}

	userID, err := result.LastInsertId()
	if err != nil {
		log.Panicln(err.Error())
	}

	return userID
}

// GetByID : returns the selected user from the database as an object
func GetByID(id int64) User {
	db, err := sql.Open("mysql", "toolset_select:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	tagRows, err := db.Query("SELECT id, email, username, creationDate FROM tbl_user WHERE id = ?", id)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer tagRows.Close()

	var u User
	var timeStr string
	for tagRows.Next() {
		err := tagRows.Scan(&u.ID, &u.Email, &u.Username, &timeStr)
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	err = tagRows.Err()
	if err != nil {
		log.Panicln(err.Error())
	}

	if u.ID == 0 && u.Email == "" && u.Username == "" {
		// when there is no entry found, return id = -1
		u.ID = -1
	}

	database.ConvertTime(&u.CreationDate, &timeStr)
	return u
}

// GetByEmail : returns the selected user from the database as an object
func GetByEmail(e int64) User {
	db, err := sql.Open("mysql", "toolset_select:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	tagRows, err := db.Query("SELECT id, email, username, creationDate FROM tbl_user WHERE email = ?", e)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer tagRows.Close()

	var u User
	var timeStr string
	for tagRows.Next() {
		err := tagRows.Scan(&u.ID, &u.Email, &u.Username, &timeStr)
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	err = tagRows.Err()
	if err != nil {
		log.Panicln(err.Error())
	}

	if u.ID == 0 && u.Email == "" && u.Username == "" {
		// when there is no entry found, return id = -1
		u.ID = -1
	}

	database.ConvertTime(&u.CreationDate, &timeStr)
	return u
}

// GetByUsername : returns the selected user from the database as an object
func GetByUsername(n int64) User {
	db, err := sql.Open("mysql", "toolset_select:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	tagRows, err := db.Query("SELECT id, email, username, creationDate FROM tbl_user WHERE email = ?", n)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer tagRows.Close()

	var u User
	var timeStr string
	for tagRows.Next() {
		err := tagRows.Scan(&u.ID, &u.Email, &u.Username, &timeStr)
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	err = tagRows.Err()
	if err != nil {
		log.Panicln(err.Error())
	}

	if u.ID == 0 && u.Email == "" && u.Username == "" {
		// when there is no entry found, return id = -1
		u.ID = -1
	}

	database.ConvertTime(&u.CreationDate, &timeStr)
	return u
}
