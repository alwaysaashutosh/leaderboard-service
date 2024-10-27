/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/alwaysaashutosh/leaderboard-service/pkg/config"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/database"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/logger"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/server"
	"github.com/alwaysaashutosh/leaderboard-service/pkg/utils/constants"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

}

func main() {

	cfg := config.Default()
	// cmd represents the base command when called without any subcommands
	var cmd = &cobra.Command{
		Use:   "Leaderboard-Service",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			startServerandRun(cfg)
		},
	}

	cmd.PersistentFlags().StringVar(&cfg.ConfigFilePath, "config", constants.DefaultConfigFile, "path to configuration file")

	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func startServerandRun(cfg *config.Config) {
	cfg.Init()
	logger.Setup(cfg.LoggerConfig)

	err := database.NewClient(cfg.DatabaseConfig)
	if err != nil {
		log.Fatal().Msgf("Failed to connect to database: %v", err)
	}

	server.Setup(cfg.ServerConfig).ServeAsync()

}
