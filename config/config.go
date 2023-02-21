package config

import (
	"flag"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type (
	// Config представляет собой набор различных конфигураций
	Config struct {
		Postgres Postgres `yaml:"postgres"`
		Logger   Logger   `yaml:"logger"`
	}

	// Postgres представляет собой конфигурацию базы данных PostgreSQL
	Postgres struct {
		User     string `yaml:"user" env:"POSTGRES_USER"`
		Password string `env:"POSTGRES_PASSWORD"`
		Host     string `yaml:"host" env:"POSTGRES_HOST"`
		Port     string `yaml:"port" env:"POSTGRES_PORT"`
		DB       string `yaml:"db" env:"POSTGRES_DB"`
		SSLMode  string `yaml:"ssl_mode" env:"POSTGRES_SSL_MODE"`
	}

	// Logger представляет собой конфигурацию логгера
	Logger struct {
		Development bool   `yaml:"development"`
		Level       string `yaml:"level"`
		Encoding    string `yaml:"encoding"`
	}
)

var configPath = flag.String("c", "config.yaml", "Path to configuration file")

// NewConfig читает конфигурации из файла и переменных окружения
func NewConfig() (*Config, error) {
	if configPath == nil {
		return nil, errors.New("invalid config path")
	}

	config := new(Config)
	if err := cleanenv.ReadConfig(*configPath, config); err != nil {
		return nil, errors.WithMessage(err, "read config")
	}

	return config, nil
}
