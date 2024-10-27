package config

import (
	"github.com/alwaysaashutosh/leaderboard-service/pkg/server"
	"github.com/spf13/viper"
)

func loadServerConfig() *server.Config {
	return &server.Config{
		Port:     viper.GetUint16("server.port"),
		BasePath: viper.GetString("server.basePath"),
		Mode:     viper.GetString("server.mode"),
	}
}
