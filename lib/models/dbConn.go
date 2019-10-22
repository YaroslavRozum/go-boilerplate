package models

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/YaroslavRozum/go-boilerplate/settings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func InitDB(settings settings.Settings) (*sqlx.DB, error) {
	var err error
	connectionString := fmt.Sprintf(
		`user=%s dbname=%s sslmode=%s port=%s`,
		settings.DbUser,
		settings.DbName,
		settings.DbSslMode,
		settings.DbPort,
	)
	db, err = sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Mappers struct {
	UserMapper    *UserMapper
	ProductMapper *ProductMapper
}

func CreateMappers(
	queryBuilder squirrel.StatementBuilderType,
	db *sqlx.DB,
) Mappers {
	return Mappers{
		UserMapper:    newUserMapper(QueryBuilder, db),
		ProductMapper: newProductMapper(QueryBuilder, db),
	}
}
