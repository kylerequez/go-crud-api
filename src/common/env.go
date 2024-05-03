package common

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	log.Println(":::-::: Loading env variables...")
	err := godotenv.Load()
	if err != nil {
		return err
	}
	log.Println(":::-::: Successfully loaded env variables")
	return nil
}
