package repositories

import (
	"database/sql"
	"encoding/json"
	"fds/src/models"
	"io"
	"net/http"
)

type UserRepositoryinterface interface {
	Create(w http.ResponseWriter, r *http.Request, db *sql.DB) error
	GetByName(w http.ResponseWriter, r *http.Request, db *sql.DB, name string) (*models.User, error)
}

type UserRepository struct {
}

func (repository UserRepository) GetByName(w http.ResponseWriter, r *http.Request, db *sql.DB, name string) (*models.User, error) {

	user := models.User{}

	row := db.QueryRow("SELECT * FROM users where name=?", name)

	if err := row.Scan(&user.Name, &user.Age); err != nil {
		return &user, err
	}

	return &user, nil
}

func (repository UserRepository) Create(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}

	user := models.User{}

	if err := json.Unmarshal(body, &user); err != nil {
		return err
	}

	_, queryError := db.Exec("INSERT INTO users (name, age) values (?, ?)", user.Name, user.Age)

	if queryError != nil {
		return queryError
	}

	return nil
}
