package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"yogan.dev/snippetbox/pkg/models"
)

// UserModel holds a references to the DB
type UserModel struct {
	DB *sql.DB
}

// Insert adds a new user to the users table
func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (name, email, hashed_password, created)
	VALUES(?, ?, ?, UTC_TIMESTAMP())`

	_, err = m.DB.Exec(query, name, email, string(hashedPassword))
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}

	return nil
}

// Authenticate verifies the user exists and their password matches the hashed password
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get takes a user ID and returns that user from the db layer
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
