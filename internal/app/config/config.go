package config

import (
	"errors"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env    string `yaml:"env" env:"ENV" env-default:"local"`
	Server `yaml:"server"`
	DB     `yaml:"db"`
}

type Server struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port int    `yaml:"port" env-default:"8080"`
}

type DB struct {
	URL string `yaml:"URL" env-required:"true"`
}

func MustLoad() (*Config, error) {
	config := Config{}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		defaultPath := "./config/local.yaml"
		configPath = defaultPath
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New(fmt.Sprintf("config file is not exist: %s", configPath))
	}

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot read config: %s", err))
	}
	return &config, nil
}
