package mysql

import (
	"database/sql"

	"github.com/Farhankhan00/snippetbox/pkg/models"
)

// UserModel models
type UserModel struct {
	DB *sql.DB
}

// Insert insert
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate user authenticate
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get get
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
