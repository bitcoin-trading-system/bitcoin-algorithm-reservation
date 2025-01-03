package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	BaseConfig BaseConfig `toml:"baseConfig"`
	DBConfig   DBConfig
}

type BaseConfig struct {
	Port string `toml:"port"`
}

type DBConfig struct {
	Host         string
	User         string
	Password     string
	DataBaseName string
}

func NewConfig(tomlFilePath, envFilePath string) Config {
	var config Config

	if _, err := toml.DecodeFile(tomlFilePath, &config); err != nil {
		panic(err)
	}

	if err := godotenv.Load(envFilePath); err != nil {
		panic(err)
	}

	config.DBConfig.Host = os.Getenv("MYSQL_HOST")
	config.DBConfig.User = os.Getenv("MYSQL_USER")
	config.DBConfig.Password = os.Getenv("MYSQL_PASSWORD")
	config.DBConfig.DataBaseName = os.Getenv("MYSQL_DATABASE")

	if err := config.mustCheck(); err != nil {
		panic(err)
	}

	return config
}

func (cfg Config) mustCheck() error {
	if cfg.BaseConfig.Port == "" {
		return errors.New("port is empty")
	}

	if cfg.DBConfig.Host == "" {
		return errors.New("host is empty")
	}

	if cfg.DBConfig.User == "" {
		return errors.New("user is empty")
	}

	if cfg.DBConfig.Password == "" {
		return errors.New("password is empty")
	}

	if cfg.DBConfig.DataBaseName == "" {
		return errors.New("database name is empty")
	}

	return nil

}
