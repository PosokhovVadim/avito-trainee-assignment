package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServiceName string     `yaml:"service_name" env-required:"true"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server" env-required:"true"`
}

type HTTPServer struct {
	Port string `yaml:"port" env-required:"true"`
}

func New(filePath string) (*Config, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return &Config{}, err
	}
	var config Config
	err := cleanenv.ReadConfig(filePath, &config)

	return &config, err
}
