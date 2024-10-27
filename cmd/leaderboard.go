package cmd

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

func Execute() {
	cfg := config.Default()
	// cmd represents the base command when called without any subcommands
	var cmd = &cobra.Command{
		Use:   "Leaderboard-Service",
		Short: "Starts the leaderboard service to manage user ranks and submissions.",
		Long: `The Leaderboard-Service is a server application that manages user ranking data and submissions 
across various scopes. It provides endpoints for fetching user ranks, submitting user data, 
and querying top ranks, supporting efficient leaderboard management.`,
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
