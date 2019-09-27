package settings

import (
	"fmt"
	"os"
)

type Settings struct {
	JwtSecret []byte
	Port      string
	DbName    string
	DbUser    string
	DbSslMode string
	DbPort    string
}

var DefaultSettings Settings

func InitSettings() {
	DefaultSettings = Settings{
		JwtSecret: []byte("THIS_IS_VERY_BIG_SECRET"),
		Port:      ":3000",
		DbName:    os.Getenv("DB_NAME"),
		DbUser:    os.Getenv("DB_USER"),
		DbSslMode: os.Getenv("DB_SSL_MODE"),
		DbPort:    os.Getenv("DB_PORT"),
	}
	fmt.Println(DefaultSettings)
}
