package models

import (
	"database/sql"
	"time"
)

// Snippet type holds the data for an individual snippet.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert inserts a new snippet into the database.
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 0, nil
}

// Get returns specific snippet based on id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Latest returns 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
