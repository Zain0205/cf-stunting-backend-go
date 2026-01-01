package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// OPTIONAL: hanya untuk local development
	if err := godotenv.Load(); err != nil {
		log.Println("ENV loaded from system (docker or server)")
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
