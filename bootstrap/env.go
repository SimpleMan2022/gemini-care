package bootstrap

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

type ENV struct {
	DB_HOST              string
	DB_PORT              string
	DB_USER              string
	DB_PASSWORD          string
	DB_NAME              string
	DB_SSL_MODE          string
	DB_TIME_ZONE         string
	PORT                 string
	GOOGLE_CLIENT_ID     string
	GOOGLE_CLIENT_SECRET string
	GOOGLE_REDIRECT_URL  string
}

func NewEnv() *ENV {
	envPath := filepath.Join(".env")
	if err := godotenv.Load(envPath); err != nil {
		panic(err)
	}

	return &ENV{
		DB_HOST:              os.Getenv("DB_HOST"),
		DB_PORT:              os.Getenv("DB_PORT"),
		DB_USER:              os.Getenv("DB_USER"),
		DB_PASSWORD:          os.Getenv("DB_PASSWORD"),
		DB_NAME:              os.Getenv("DB_NAME"),
		DB_SSL_MODE:          os.Getenv("DB_SSL_MODE"),
		DB_TIME_ZONE:         os.Getenv("DB_TIME_ZONE"),
		PORT:                 os.Getenv("PORT"),
		GOOGLE_CLIENT_ID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GOOGLE_CLIENT_SECRET: os.Getenv("GOOGLE_CLIENT_SECRET"),
		GOOGLE_REDIRECT_URL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	}
}
