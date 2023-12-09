package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App App `yaml:"app"`
	DB  DB  `yaml:"db"`
}

type App struct {
	Port string `yaml:"port"`
}

type DB struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	Name            string `yaml:"name"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	MaxConnLifetime int    `yaml:"maxConnLifetime"`
}

var Cfg *Config

func LoadConfig(filename string) (err error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	cfg := Config{}

	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return
	}

	Cfg = &cfg
	return
}
