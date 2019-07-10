package mysql

import (
	"database/sql"
	"strings"

	"github.com/Farhankhan00/snippetbox/pkg/models"
	"github.com/go-sql-driver/mysql"
)

// UserModel models
type UserModel struct {
	DB *sql.DB
}

// Insert insert
func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GeneratePassword([]byte(password), 12)
	if err != nil {
		return nil
	}
	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES(?, ?, ?, UTC_TIMESTAMP())`
	__, err = m.DB.Exec(stmt, name, email, string(hasehedPassword))
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
	}
	return err
}

// Authenticate user authenticate
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get get
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
