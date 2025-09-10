package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Numb NumberConfig
}

type NumberConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using dafault config")
	}
	return &Config{
		Numb: NumberConfig{
			Secret: os.Getenv("SECRET"),
		},
	}
}
