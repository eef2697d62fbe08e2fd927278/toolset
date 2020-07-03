package user

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/youngtrashbag/toolset/src/database"
)

// User : a struct, so new users can be added to the site
type User struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	password     string
	CreationDate time.Time `json:"creation_date"`
}

// NewUser : returns an object of User struct
//			with username, password and already hashed password
func NewUser(e string, un string, p string) User {
	var u User

	u.Email = e
	u.Username = un
	u.password = hashPassword(p)
	u.CreationDate = time.Now()

	return u
}

// hashPassword : sets the Password of the User and automatically hashes it
func hashPassword(p string) string {
	hashed := sha256.Sum256([]byte(p))
	return hex.EncodeToString(hashed[:])
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
	insertUser, err := db.Prepare("INSERT INTO tbl_user (email, username, password, creationDate) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer insertUser.Close()

	var time string
	database.ConvertTime(&u.CreationDate, &time)
	// execute sql statement
	_, err = insertUser.Exec(u.Email, u.Username, u.password, time)
	if err != nil {
		log.Panicln(err.Error())
	}
}
