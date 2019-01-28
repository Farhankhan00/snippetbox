package mysql

import (
	"database/sql"

	"github.com/Farhankhan00/snippetbox/pkg/models"
)

// SnippetModel Defines a type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get will return a specific snippet based on Id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
