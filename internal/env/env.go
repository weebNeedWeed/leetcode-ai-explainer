package env

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Overload("./env")
}

func GetString(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	} else {
		return fallback
	}
}
