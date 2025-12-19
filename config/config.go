package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	LinkedIn struct {
		DailyConnectionLimit int `yaml:"daily_connection_limit"`
		DailyMessageLimit    int `yaml:"daily_message_limit"`
	} `yaml:"linkedin"`

	Browser struct {
		Headless bool `yaml:"headless"`
	} `yaml:"browser"`

	Delays struct {
		MinActionDelayMs int `yaml:"min_action_delay_ms"`
		MaxActionDelayMs int `yaml:"max_action_delay_ms"`
	} `yaml:"delays"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
