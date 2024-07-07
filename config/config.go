package config

import "github.com/joho/godotenv"

func EnvInit() error {
	err := godotenv.Load()

	return err
}
