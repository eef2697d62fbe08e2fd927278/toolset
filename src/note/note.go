package note

/*
import (
	"time"
)
*/

// Note : Struct used for writing note
type Note struct {
	title   string
	content string
	time    time
	tags    []string
}

// NewNote : a constructor for the Note struct
func NewNote(t string, c string, t time, tg []string) Note {
	var n Note

	n.SetTitle(t)
	n.SetContent(c)
	n.SetTime(t)
	n.SetTags(tg)

	return n
}

// SetTitle : set title of note
func (n *Note) SetTitle(t string) {
	n.title = t
}

// SetContent : set content of note
func (n *Note) SetContent(c string) {
	n.content = c
}

// SetTime : sets time of note
func (n *Note) SetTime(t time) {
	n.time = t
}

// SetTags : set tags of note
func (n *Note) SetTags(t []string) {
	n.tags = t
}
