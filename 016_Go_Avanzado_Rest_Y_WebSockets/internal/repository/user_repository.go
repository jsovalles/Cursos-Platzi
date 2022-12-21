package repository

import (
	"database/sql"
	"fmt"
	"github.com/jsovalles/rest-ws/internal/models"
	"github.com/jsovalles/rest-ws/internal/utils"
)

type UserRepository interface {
	CreateUser(user models.User) (err error)
	GetUserById(id string) (user models.User, err error)
	GetUserByEmail(email string) (user models.User, err error)
}

type userRepository struct {
	env utils.Env
	db  utils.Database
}

func NewUserRepository(env utils.Env, db utils.Database) UserRepository {
	return &userRepository{env: env, db: db}
}

const (
	usersTable     = "users"
	createUser     = "INSERT INTO " + usersTable + " (id, email, password) VALUES (:id, :email, :password)"
	getUserById    = "SELECT * FROM " + usersTable + " WHERE id = $1"
	getUserByEmail = "SELECT * FROM " + usersTable + " WHERE email = $1"
)

func (u *userRepository) CreateUser(user models.User) (err error) {
	_, err = u.db.Client.NamedExec(createUser, user)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	return
}

func (u *userRepository) GetUserById(id string) (user models.User, err error) {
	err = u.db.Client.Get(&user, getUserById, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("there are no results for this user, please validate")
			return
		}
		return
	}
	return
}

func (u *userRepository) GetUserByEmail(email string) (user models.User, err error) {
	err = u.db.Client.Get(&user, getUserByEmail, email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("there are no results for this user, please validate")
			return
		}
		return
	}
	return
}
