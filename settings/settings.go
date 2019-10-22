package settings

import (
	"os"
)

type Settings struct {
	JwtSecret      []byte
	Port           string
	DbName         string
	DbUser         string
	DbSslMode      string
	DbPort         string
	SenderEmail    string
	SenderPassword string
}

func CreateSettings() Settings {
	return Settings{
		JwtSecret:      []byte(os.Getenv("JWT_SECRET")),
		Port:           os.Getenv("PORT"),
		DbName:         os.Getenv("DB_NAME"),
		DbUser:         os.Getenv("DB_USER"),
		DbSslMode:      os.Getenv("DB_SSL_MODE"),
		DbPort:         os.Getenv("DB_PORT"),
		SenderEmail:    os.Getenv("SENDER_EMAIL"),
		SenderPassword: os.Getenv("SENDER_PASSWORD"),
	}
}
