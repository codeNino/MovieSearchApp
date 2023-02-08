package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var (
	DatabaseURI string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=verify-full",
		goDotEnvVariable("DBHost"), goDotEnvVariable("DBUser"), goDotEnvVariable("DBPwd"), goDotEnvVariable("DBName"))
	OMDB_API_KEY  string = goDotEnvVariable("omdbAPIKey")
	Client_Domain string = goDotEnvVariable("clientDomain")
)
