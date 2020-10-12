package models

import (
	"errors"
	"time"
)

// ErrNoRecord returns a new error on no record found
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet defines the structure of a snippet model
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
