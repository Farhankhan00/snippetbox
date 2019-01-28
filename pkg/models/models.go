package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: No matching records found")

//Snippet struct provides data for snippetbox snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
