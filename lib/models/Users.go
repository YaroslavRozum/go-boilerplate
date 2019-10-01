package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var DefaultUserMapper *UserMapper

func initUserMapper() {
	DefaultUserMapper = &UserMapper{}
}

type User struct {
	ID       string `json:"id" db:"id"`
	UserName string `json:"username" db:"username"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Surname  string `json:"surname" db:"surname"`
	Password string `json:"password,omitempty" db:"password"`
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

	query, args, err := QueryBuilder.
		Insert("users").
		Columns("id", "username", "name", "email", "surname", "password").
		Values(user.ID, user.UserName, user.Name, user.Email, user.Surname, user.Password).
		ToSql()

	_, err = DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uM *UserMapper) FindOne(where interface{}, args ...interface{}) (*User, error) {
	query, queryArgs, _ := QueryBuilder.
		Select("*").
		From("users").
		Where(where, args...).
		Limit(1).
		ToSql()
	row := DB.QueryRowx(query, queryArgs...)
	user := &User{}
	row.StructScan(user)

	return user, nil
}

func (uM *UserMapper) FindAll(where interface{}, limit, offset uint64, args ...interface{}) ([]*User, error) {
	selectBuilder := QueryBuilder.
		Select("*").
		From("users").
		Where(where, args...)
	if limit != 0 {
		selectBuilder = selectBuilder.Limit(limit)
	}
	if offset != 0 {
		selectBuilder = selectBuilder.Offset(offset)
	}

	query, queryArgs, _ := selectBuilder.ToSql()

	rows, err := DB.Queryx(query, queryArgs...)
	if err != nil {
		return nil, err
	}

	users := initUsersByLimit(limit)
	for rows.Next() {
		user := &User{}
		err := rows.StructScan(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func initUsersByLimit(limit uint64) []*User {
	if limit != 0 {
		return make([]*User, 0, limit)
	}
	return []*User{}
}
