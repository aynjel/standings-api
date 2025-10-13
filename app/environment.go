package app

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		// don't fatal here - allow running without a .env file (useful for simple dev runs)
		log.Println("warning: .env file not found, continuing with environment variables")
	}
}
