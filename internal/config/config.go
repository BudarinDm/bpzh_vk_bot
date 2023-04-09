package config

import (
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	App AppConfig
	DB  DBConfig
	Api ApiConfig
}

type AppConfig struct {
	Port     string
	LogLevel string
	BotToken string
}

type DBConfig struct {
	DBString string
	FSConf   string
}

type ApiConfig struct {
	Url   string
	Token string
}

func ReadConfig() (*Config, error) {

	var config Config
	var err error

	//app parse

	config.App.Port = os.Getenv("SERVER_PORT")
	if config.App.Port == "" {
		config.App.Port = "80"
	}

	config.App.LogLevel = os.Getenv("LOG_LEVEL")
	if config.App.LogLevel == "" {
		config.App.LogLevel = "debug"
	}

	config.Api.Url = os.Getenv("API_URL")
	if config.Api.Url == "" {
		return nil, errors.New("Not specified API_URL")
	}

	config.Api.Token = os.Getenv("API_TOKEN")
	if config.Api.Token == "" {
		return nil, errors.New("Not specified API_TOKEN")
	}

	config.App.BotToken = os.Getenv("BOT_TOKEN")
	if config.App.BotToken == "" {
		return nil, errors.New("Not specified BOT_TOKEN")
	}

	config.DB.FSConf = os.Getenv("FS_CONF")
	if config.DB.FSConf == "" {
		return nil, errors.New("Not specified FS_CONF")
	}

	return &config, err

}
