package models

import (
	"fmt"

	"github.com/YaroslavRozum/go-boilerplate/settings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() error {
	var err error
	defaultSettings := settings.DefaultSettings
	connectionString := fmt.Sprintf(
		`user=%s dbname=%s sslmode=%s port=%s`,
		defaultSettings.DbUser,
		defaultSettings.DbName,
		defaultSettings.DbSslMode,
		defaultSettings.DbPort,
	)
	DB, err = sqlx.Connect("postgres", connectionString)
	if err != nil {
		return err
	}
	return nil
}

func InitModels() {
	initUserMapper()
}
