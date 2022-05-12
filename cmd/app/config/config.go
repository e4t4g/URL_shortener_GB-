package config

import (
	"go.uber.org/zap"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port     int      `yaml:"port" envconfig:"port"`
	DBconfig DBconfig `yaml:"db_config" envconfig:"db_config"`
	Host     string   `yaml:"host" envconfig:"host"`
}

type DBconfig struct {
	DBurl string `yaml:"db_url" envconfig:"db_url"`
}

func (c *Config) ReadFromFile(logger *zap.SugaredLogger) error {
	configPath := "/home/e4t4g/Desktop/URLshortener/URL_shortener_GB-/cmd/configs/app.yaml"

	data, err := os.ReadFile(configPath)
	if err != nil {
		logger.Errorf("incorrect config file: %s", err)
	}

	if err = yaml.Unmarshal(data, c); err != nil {
		logger.Errorf("unnable to unmarshal file: %s", err)
	}
	return nil
}
