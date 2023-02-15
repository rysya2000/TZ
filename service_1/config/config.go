package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
