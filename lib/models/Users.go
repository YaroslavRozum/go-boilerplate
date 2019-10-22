package models

import (
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func newUserMapper(
	queryBuilder squirrel.StatementBuilderType,
	db *sqlx.DB,
) *UserMapper {
	return &UserMapper{queryBuilder, db}
}

type User struct {
	ID       string `json:"id" db:"id"`
	UserName string `json:"username" db:"username"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Surname  string `json:"surname" db:"surname"`
	Password string `json:"password,omitempty" db:"password"`
}

type UserMapper struct {
	queryBuilder squirrel.StatementBuilderType
	db           *sqlx.DB
}

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

	query, args, err := uM.queryBuilder.
		Insert("users").
		Columns("id", "username", "name", "email", "surname", "password").
		Values(user.ID, user.UserName, user.Name, user.Email, user.Surname, user.Password).
		ToSql()

	_, err = uM.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uM *UserMapper) FindOne(where interface{}, args ...interface{}) (*User, error) {
	query, queryArgs, _ := uM.queryBuilder.
		Select("*").
		From("users").
		Where(where, args...).
		Limit(1).
		ToSql()
	row := uM.db.QueryRowx(query, queryArgs...)
	user := &User{}
	row.StructScan(user)

	return user, nil
}

func (uM *UserMapper) FindAll(where interface{}, limit, offset uint64, args ...interface{}) ([]*User, error) {
	selectBuilder := uM.queryBuilder.
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

	rows, err := uM.db.Queryx(query, queryArgs...)
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
