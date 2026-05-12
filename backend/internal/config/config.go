package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DBConfig `mapstructure:"database"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type DBConfig struct {
	URL string `mapstructure:"url"`
}

var AppConfig Config

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 0
	}

	AppConfig = Config{
		Server:   ServerConfig{Port: port},
		Database: DBConfig{URL: os.Getenv("DB_URL")},
	}

	return &AppConfig, nil

}
