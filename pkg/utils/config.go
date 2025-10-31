package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port int
		Env  string
	}
	MySQL struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string `mapstructure:"dbname"`
	}
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	v.AddConfigPath("../..")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func MustLoadConfig(path string) *Config {
	c, err := LoadConfig(path)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return c
}
