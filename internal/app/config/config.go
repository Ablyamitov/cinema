package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
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
		//defaultPath := "C:/Users/ToTheMoon/Documents/go-project/cinema/config/local.yaml"
		configPath = defaultPath
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file is not exist: %s", configPath)
	}

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		return nil, fmt.Errorf("cannot read config: %s", err)
	}
	return &config, nil
}
