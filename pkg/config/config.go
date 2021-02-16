package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	dbConfig         DBConfig
	logConfig        LogConfig
	logFileConfig    LogFileConfig
	httpServerConfig HTTPServerConfig
}

func (config Config) GetDBConfig() DBConfig {
	return config.dbConfig
}

func (config Config) GetLogConfig() LogConfig {
	return config.logConfig
}

func (config Config) GetHTTPServerConfig() HTTPServerConfig {
	return config.httpServerConfig
}

func (config Config) GetLogFileConfig() LogFileConfig {
	return config.logFileConfig
}

func NewConfig(configFile string) Config {
	viper.AutomaticEnv()

	if configFile != "" {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
	}

	return Config{
		dbConfig:         newDBConfig(),
		logConfig:        newLogConfig(),
		logFileConfig:    newLogFileConfig(),
		httpServerConfig: newHTTPServerConfig(),
	}
}
