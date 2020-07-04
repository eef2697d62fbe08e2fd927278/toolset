package note

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // this is needed for mysql
)

// Note : Struct used for writing note
type Note struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreationDate time.Time `json:"datetime"`
	AuthorID     int64     `json:"author_id"`
}

// NewNote : a constructor for the Note struct
func NewNote(t string, c string) Note {
	var n Note

	n.Title = t
	n.Content = c
	n.CreationDate = time.Now()

	// TODO: set authorId

	return n
}
