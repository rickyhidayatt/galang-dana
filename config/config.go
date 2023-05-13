package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func ReloadEnv() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return err
	}
	return nil
}
