package user

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// User : a struct, so new users can be added to the site
type User struct {
	Username string
	Email    string
	password string
}

// NewUser : returns an object of User struct
//			with username, password and already hashed password
func NewUser(e string, un string, p string) User {
	var u User

	u.Email = e
	u.Username = un
	u.SetPassword(p)

	return u
}

// SetPassword : sets the Password of the User and automatically hashes it
func (u *User) SetPassword(p string) {
	hashed := sha256.Sum256([]byte(p))
	u.password = hex.EncodeToString(hashed[:])
}

// Insert : saves a user in the database
func (u *User) Insert() {
	// connection to database
	db, err := sql.Open("mysql", "toolset_insert:password@/toolset")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer db.Close()

	// prepare sql statement
	insertUser, err := db.Prepare("INSERT INTO tbl_user (email, username, password) VALUES (?, ?, ?)")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer insertUser.Close()

	// execute sql statement
	_, err = insertUser.Exec(u.Email, u.Username, u.password)
	if err != nil {
		log.Panicln(err.Error())
	}
}
