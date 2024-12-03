package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"5050"`
	} `yaml:"server"`
	Database struct {
		URL string `yaml:"url"`
	} `yaml:"database"`
}

func InitConfig(path string) (*Config, error) {
	config := new(Config)

	if err := cleanenv.ReadConfig(path, config); err != nil {
		log.Fatalf("Error loading config: %v", err)
		return nil, err
	}
	return config, nil
}
