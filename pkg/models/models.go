package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: No matching records found")

	ErrInvalidCredentials = errors.New("models: invalid credentias")

	ErrDuplicateEmail = errors.New("models: duplicate email")
)

//Snippet struct provides data for snippetbox snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User models user var snippetbox
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}
