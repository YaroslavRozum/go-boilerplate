package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var DefaultUserMapper *UserMapper

type User struct {
	ID       string `db:"id"`
	UserName string `db:"username"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Surname  string `db:"surname"`
	Password string `db:"password"`
}

type UserMapper struct{}

func (uM *UserMapper) NewUser(username, name, email, surname, password string) (*User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	user := &User{
		ID:       id,
		Password: string(pass),
		UserName: username,
		Email:    email,
		Name:     name,
		Surname:  surname,
	}

	query, args, _ := QueryBuilder.
		Insert("users").
		Columns("id", "username", "name", "email", "surname", "password").
		Values(user.ID, user.UserName, user.Name, user.Email, user.Surname, user.Password).
		ToSql()

	_, err = Conn.Exec(query, args)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func initUserMapper() {
	DefaultUserMapper = &UserMapper{}
}
