package config

import (
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	App AppConfig
	DB  DBConfig
}

type AppConfig struct {
	Port     string
	LogLevel string
	BotToken string
}

type DBConfig struct {
	DBString string
	FSPath   string
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

	config.App.BotToken = os.Getenv("BOT_TOKEN")
	if config.App.BotToken == "" {
		return nil, errors.New("Not specified BOT_TOKEN")
	}

	config.DB.FSPath = os.Getenv("FS_PATH")
	if config.DB.FSPath == "" {
		return nil, errors.New("Not specified FS_PATH")
	}

	//db parse

	//config.DB.DBString = os.Getenv("DBSTRING")
	//if config.DB.DBString == "" {
	//	return nil, errors.New("Not specified DBSTRING")
	//}

	return &config, err

}
