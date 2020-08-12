package user

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // this is needed for mysql
	"github.com/youngtrashbag/toolset/pkg/utils"
)

// Insert : saves a user in the database
func (u *User) Insert() int64 {

	TODO: do i really need validation here, or do i handle that in frontend ?!?!

	// validate, if -1 then there is no entry in db -> continue
	if GetByEmail(u.Email).ID != -1 {
		return -2
	}
	if GetByUsername(u.Username).ID != -1 {
		return -2
	}

	// connection to database
	db, err := sql.Open("mysql", util.CNInsert)
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
	utils.ConvertTime(&u.CreationDate, &time)

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

// GetByID : returns the selected user from the utils as an object
func GetByID(id int64) User {
	db, err := sql.Open("mysql", utils.CNSelect)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	userRows, err := db.Query("SELECT id, email, username, creationDate FROM tbl_user WHERE id = ?", id)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer userRows.Close()

	var u User
	var timeStr string
	for userRows.Next() {
		err := userRows.Scan(&u.ID, &u.Email, &u.Username, &timeStr)
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	if userRows.Err() != nil {
		log.Panicln(userRows.Err())
	}

	if u.ID == 0 && u.Email == "" && u.Username == "" {
		// when there is no entry found, return id = -1
		u.ID = -1
	}

	utils.ConvertTime(&u.CreationDate, &timeStr)
	return u
}

//TODO: is this neede, because of privacy concerns

// GetByEmail : returns the selected user from the database as an object
func GetByEmail(e string) User {
	db, err := sql.Open("mysql", utils.CNSelect)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	userRows, err := db.Query("SELECT id, email, username, creationDate FROM tbl_user WHERE email = ?", e)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer userRows.Close()

	var u User
	var timeStr string
	for userRows.Next() {
		err := userRows.Scan(&u.ID, &u.Email, &u.Username, &timeStr)
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	if userRows.Err() != nil {
		log.Panicln(userRows.Err())
	}

	if u.ID == 0 && u.Email == "" && u.Username == "" {
		// when there is no entry found, return id = -1
		u.ID = -1
	}

	utils.ConvertTime(&u.CreationDate, &timeStr)
	return u
}

// GetByUsername : returns the selected user from the database as an object
func GetByUsername(n string) User {
	db, err := sql.Open("mysql", utils.CNSelect)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	userRows, err := db.Query("SELECT id, email, username, creationDate FROM tbl_user WHERE email = ?", n)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer userRows.Close()

	var u User
	var timeStr string
	for userRows.Next() {
		err := userRows.Scan(&u.ID, &u.Email, &u.Username, &timeStr)
		if err != nil {
			log.Panicln(err.Error())
		}
	}

	if userRows.Err() != nil {
		log.Panicln(userRows.Err())
	}

	if u.ID == 0 && u.Email == "" && u.Username == "" {
		// when there is no entry found, return id = -1
		u.ID = -1
	}

	utils.ConvertTime(&u.CreationDate, &timeStr)
	return u
}

