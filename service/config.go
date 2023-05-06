package service

import (
	"gopkg.in/yaml.v3"
	"os"
)

type AsyncTasks struct {
	ChannelSize  int `yaml:"channel_size"`
	WorkersCount int `yaml:"workers_count"`
}

type Config struct {
	Port       int        `yaml:"port"`
	Tasks      AsyncTasks `yaml:"tasks"`
	HttpClient AsyncTasks `yaml:"http_client"`
}

func GetConfig() (*Config, error) {
	yamlFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
