package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/linusfri/calc-api/db"
	"github.com/linusfri/calc-api/routes"
)

func main() {
	godotenv.Load()
	checkEnvironment()
	db.Migrate()
	routes.InitRoutes()
}

func checkEnvironment() {
	requiredEnvironmentVariables := []string{
		"DB_PW",
		"DB_USER",
		"DB_SERVICE",
		"DB_PORT",
		"DB_NAME",
		"AUTH_SECRET_KEY",
		"ALLOWED_ORIGINS",
	}

	for _, envVariable := range requiredEnvironmentVariables {
		_, envVariableExists := os.LookupEnv(envVariable)

		if !envVariableExists {
			panic(fmt.Sprintf(
				"ERROR: You have to define the %s variable in either your server environment or .env file",
				envVariable,
			))
		}
	}
}
