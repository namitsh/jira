package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading configurations from env file: %s", err.Error())
	}
	log.Println("Successfully loaded the environment variables")
	return os.Getenv("MONGODB_URI")
}
