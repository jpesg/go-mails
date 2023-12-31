package shared

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file")
	}

	return os.Getenv(key)
}
