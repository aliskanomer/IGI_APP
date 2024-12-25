package utils

import (
	// modules
	"log"
	"os"

	// third-party
	"github.com/joho/godotenv"
)

func LoadEnv() {
	/*
		// read env file
		err := godotenv.Load()

		// on error; terminate application and log error
		if err != nil {
			log.Fatal("ERR: Server configuration failed. Please check environment files!")
		}

		// log success and continue
		Logger("success", "Server", 0, "configured!")*/

	// Only load .env in non-Docker environments
	if os.Getenv("DOCKER_ENV") == "" {
		err := godotenv.Load()
		if err != nil {
			Logger("error", "ServerCONFG", 0, "Failed to load environment file!")
		}
		Logger("info", "ServerCONFG", 0, "Environment file loaded!")
	} else {
		Logger("info", "ServerCONFG", 0, "Using Docker Provided Environment file!")
	}
}

func ReadEnvVar(key string) string {
	// read environment variable
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("ERR: Can not found %s in environment files", key)
	}
	return value
}
