package config

import (
    "log"
    "os"
    "gopkg.in/yaml.v3"
)

type Config struct {
	Server Server `yaml:"server"`
	Postgres Postgres `yaml:"postgres"`
	Logging Logging `yaml:"logging"`
}

type Postgres struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"root"`
	DBName   string `yaml:"dbname" env-default:""`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Logging struct {
	Level string `yaml:"level"`
}

func LoadConfig() *Config {
    var cfg Config


    file, err := os.Open("../../config/config.yaml")
    if err != nil {
        log.Fatalf("Failed to open config file: %v", err)
    }
    defer file.Close()

    // Decode the YAML file into the Config struct
    decoder := yaml.NewDecoder(file)
    if err := decoder.Decode(&cfg); err != nil {
        log.Fatalf("Failed to decode config file: %v", err)
    }

    return &cfg
}