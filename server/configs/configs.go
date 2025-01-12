package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(variable string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	v := os.Getenv(variable)
	if v == "" {
		panic("Undefined environment variable: " + variable)
	} else {
		return v
	}
}
