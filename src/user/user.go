package user

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql" // this is needed for mysql
)

// User : a struct, so new users can be added to the site
type User struct {
	ID           int64 `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	password     string
	CreationDate time.Time `json:"creation_date"`
}

// NewUser : returns an object of User struct
//			with username, password and already hashed password
func NewUser(e string, un string, p string) User {
	var u User

	u.Email = strings.ToLower(e)
	u.Username = strings.ToLower(un)
	u.password = hashPassword(p)
	u.CreationDate = time.Now()

	return u
}

// hashPassword : sets the Password of the User and automatically hashes it
func hashPassword(p string) string {
	hashed := sha256.Sum256([]byte(p))
	return hex.EncodeToString(hashed[:])
}
