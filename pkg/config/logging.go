package config

import (
	"github.com/alwaysaashutosh/leaderboard-service/pkg/logger"
	"github.com/spf13/viper"
)

func loadLogConfig() *logger.Config {
	return &logger.Config{
		Level: viper.GetString("logging.level"),
	}
}
