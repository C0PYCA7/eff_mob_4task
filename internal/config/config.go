package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env      string         `yaml:"env"`
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func Load() *Config {
	cfgPath := "config/config.yaml"

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatal("cfg doesn't exists ", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Fatal("failed to read cfg ", err)
	}

	return &cfg
}
