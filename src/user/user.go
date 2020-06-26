package user

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/youngtrashbag/toolset/src/database"

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
func NewUser(un string, e string, p string) User {
	var u User

	u.Username = un
	u.Email = e
	u.SetPassword(p)

	return u
}

// SetPassword : sets the Password of the User and automatically hashes it
func (u User) SetPassword(p string) {
	hashed := sha256.Sum256([]byte(p))
	u.password = hex.EncodeToString(hashed[:])
}

// Insert : saves a user in the database
func (u User) Insert() {

	// connection to database TODO: move this to seperate file (database.go) so everything is organized
	db := database.Connect()
	defer db.Close()

	// prepare sql statement
	insertUser, err := db.Prepare("INSERT INTO tbl_user (email, username, password) VALUES (?, ?, ?")
	if err != nil {
		panic(err.Error())
	}
	defer insertUser.Close()

	// execute sql statement
	_, err = insertUser.Exec(u.Email, u.Username, u.password)
	if err != nil {
		panic(err.Error())
	}
}
