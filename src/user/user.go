package user

import (
	"crypto/sha256"
	"encoding/hex"
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
