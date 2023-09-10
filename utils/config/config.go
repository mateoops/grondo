package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type JobConfig struct {
	Name    string `yaml:"name"`
	Cron    string `yaml:"cron"`
	Enabled bool   `yaml:"enabled"`
}

type ServerConfig struct {
	Port        string `yaml:"port"`
	ContextPath string `yaml:"contextPath"`
	ApiVersion  string `yaml:"apiVersion"`
}

type YamlConfig struct {
	Server ServerConfig `yaml:"server"`
	Jobs   []JobConfig  `yaml:"jobs"`
}

var (
	configOnce sync.Once
	AppConfig  *YamlConfig
)

func LoadYamlConfig() (*YamlConfig, error) {
	configOnce.Do(func() {
		data, err := os.ReadFile("config.yaml")
		if err != nil {
			// Handle the error as needed
			panic(err)
		}

		var config YamlConfig
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			// Handle the error as needed
			panic(err)
		}

		AppConfig = &config
	})

	return AppConfig, nil
}
