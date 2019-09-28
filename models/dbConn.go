package models

import (
	"fmt"

	"github.com/YaroslavRozum/go-boilerplate/settings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Conn *sqlx.DB

func InitConn() error {
	var err error
	defaultSettings := settings.DefaultSettings
	connectionString := fmt.Sprintf(
		`user=%s dbname=%s sslmode=%s port=%s`,
		defaultSettings.DbUser,
		defaultSettings.DbName,
		defaultSettings.DbSslMode,
		defaultSettings.DbPort,
	)
	Conn, err = sqlx.Connect("postgres", connectionString)
	if err != nil {
		return err
	}
	return nil
}
