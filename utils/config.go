package utils

import (
	"log"
	"os"

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
}

type YamlConfig struct {
	Server ServerConfig `yaml:"server"`
	Jobs   []JobConfig  `yaml:"jobs"`
}

// ToJSON returns a JSON string of the config
func readConfig() YamlConfig {
	yamlData, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var config YamlConfig

	if err := yaml.Unmarshal(yamlData, &config); err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}

// Config is the global config object
var Config YamlConfig = readConfig()
