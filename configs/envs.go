package configs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func ternaryMap(value string, defaultValue string) string {
	return map[bool]string{true: value, false: defaultValue}[len(value) != 0]
}

func getEnv(envName string, defaultValue string) string {
	return ternaryMap(os.Getenv(envName), defaultValue)
}

// Envs has values for environment variables and the defaults for them
var Envs = map[string]string{
	"IS_ENABLE_SWAGGER": getEnv("IS_ENABLE_SWAGGER", "true"),
	"IS_SERVERLESS":     getEnv("IS_SERVERLESS", "false"),
	"SERVER_PORT":       getEnv("PORT", "8080"),
	"PSQL_URI":          getEnv("PSQL_URI", "host=localhost user=postgres password=1234 dbname=database port=5432 sslmode=disable TimeZone=Asia/Seoul"),
}
