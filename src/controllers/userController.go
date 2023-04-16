package controllers

import (
	"database/sql"
	"fds/src/models"
	"fds/src/repositories"
	"net/http"
)

type UserController struct {
	userRepository repositories.UserRepositoryinterface
}

func New(userRepository repositories.UserRepositoryinterface) UserController {
	return UserController{userRepository: userRepository}
}

func (controller *UserController) GetByName(w http.ResponseWriter, r *http.Request, db *sql.DB, name string) (*models.User, error) {
	user, err := controller.userRepository.GetByName(w, r, db, name)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (controller *UserController) Store(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	err := controller.userRepository.Create(w, r, db)

	if err != nil {
		return err
	}

	return nil
}
