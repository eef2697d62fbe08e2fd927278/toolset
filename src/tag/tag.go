package tag

import (
	"time"
)

// Tag : struct for tags, which are used to organise notes
type Tag struct {
	ID           int64
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creation_date"`
}

// NewTag : returns a tag object
func NewTag(n string) Tag {
	var t Tag
	t.Name = n
	t.CreationDate = time.Now()

	return t
}
