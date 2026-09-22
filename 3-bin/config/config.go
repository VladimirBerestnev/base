package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалось найти .env файл")
	}
	key := os.Getenv("KEY")
	return &Config{
		Key: key,
	}
}
