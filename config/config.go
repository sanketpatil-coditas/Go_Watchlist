package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type PostgresConfig struct {
	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"postgres"`
}

var AppConfig PostgresConfig

func LoadConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}
}
