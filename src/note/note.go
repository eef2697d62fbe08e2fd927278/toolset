package note

import (
	"time"
)

// Note : Struct used for writing note
type Note struct {
	Title   string
	Content string
	time    time.Time
	Tags    []string
}

// NewNote : a constructor for the Note struct
func NewNote(t string, c string, tg []string) Note {
	var n Note

	n.Title = t
	n.Content = c
	n.SetTime(time.Now())
	n.Tags = tg

	return n
}

// SetTime : sets time of note
func (n *Note) SetTime(t time.Time) {
	n.time = t
}
