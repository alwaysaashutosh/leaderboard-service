package config

import (
	"github.com/alwaysaashutosh/leaderboard-service/pkg/database"
	"github.com/spf13/viper"
)

func loadDatabaseConfig() *database.Config {
	config := database.Config{
		Host:        viper.GetString("database.host"),
		Port:        viper.GetString("database.port"),
		Name:        viper.GetString("database.name"),
		IdleConn:    viper.GetInt("database.idleConn"),
		OpenConn:    viper.GetInt("database.openConn"),
		Username:    viper.GetString("database.username"),
		Password:    viper.GetString("database.password"),
		AutoMigrate: viper.GetBool("database.autoMigrate"),
	}

	return &config
}
