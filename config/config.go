package config

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var GET *Config = nil

type Config struct {
	Server struct {
		Database struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
		} `yaml:"database"`
		Web struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"web"`
	} `yaml:"server"`
	Client struct {
		EnableWizz bool `yaml:"enable_wizz"`
	} `yaml:"client"`
}

func Load() error {
	filepath := "config.yml"

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		filepath = "/etc/cao/config.yml"
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return errors.New("could not find config file")
	}

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	cfg := Config{}
	err = yaml.Unmarshal(data, &cfg)

	GET = &cfg

	return err

}
