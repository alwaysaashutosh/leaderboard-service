package config

import (
	"fmt"
	"strings"

	"github.com/alwaysaashutosh/leaderboard-service/pkg/database"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/logger"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	ConfigFilePath string
	ServerConfig   *server.Config
	DatabaseConfig *database.Config
	LoggerConfig   *logger.Config
}

func Default() *Config {
	return &Config{}
}

func (cfg *Config) read() *Config {

	viper.SetConfigFile(cfg.ConfigFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Msgf("Reading config file[%s] failed. %s", viper.ConfigFileUsed(), err)
	}

	return cfg
}

func (cfg *Config) print() *Config {
	allKeys := viper.AllKeys()
	configArr := []string{}

	for _, key := range allKeys {
		value := viper.Get(key)
		configArr = append(configArr, fmt.Sprintf("%s=%v", key, value))
	}

	log.Info().Msgf("Application configuration - [%s]", strings.Join(configArr, ", "))

	return cfg
}

func (cfg *Config) load() *Config {
	cfg.DatabaseConfig = loadDatabaseConfig()
	cfg.ServerConfig = loadServerConfig()
	cfg.LoggerConfig = loadLogConfig()

	return cfg
}

func (cfg *Config) Init() {
	// Load application config
	cfg.read().print().load()
}
