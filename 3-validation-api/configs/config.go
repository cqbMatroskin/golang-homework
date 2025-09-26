package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db DbConfig
	Auth AuthConfig
	EmailVerify EmailVerifyConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type EmailVerifyConfig struct {
	Email string
	Password string
	Address string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
		EmailVerify: EmailVerifyConfig{
			Email: os.Getenv("EMAIL"),
			Password: os.Getenv("PASSWORD"),
			Address: os.Getenv("ADDRESS"),
		},
	}
}