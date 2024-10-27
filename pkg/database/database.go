package database

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host        string
	Port        string
	Name        string
	Username    string
	Password    string
	IdleConn    int
	OpenConn    int
	AutoMigrate bool
	LogLevel    string
}

var (
	err    error
	gormDB *gorm.DB
)

func NewClient(cfg *Config) error {
	if cfg == nil {
		log.Fatal().Msgf("database config not provided")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=UTC", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)

	level := parseLogLevel(cfg.LogLevel)
	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(level),
	})
	if err != nil {
		return fmt.Errorf("cannot connect to database at %s:%v/%s due to %v", cfg.Host, cfg.Port, cfg.Name, err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("cannot get sql db instance due to %v", err)
	}

	sqlDB.SetMaxIdleConns(cfg.IdleConn)
	sqlDB.SetMaxOpenConns(cfg.OpenConn)

	err = sqlDB.Ping()
	if err != nil {
		return fmt.Errorf("pinging the database failed with %v", err)
	}
	if cfg.AutoMigrate {
		gormDB.AutoMigrate(&Leaderboard{})
	}

	log.Info().Msgf("Successfully connected to database at %s:%v/%s", cfg.Host, cfg.Port, cfg.Name)

	return err
}

func Handler() *gorm.DB {
	return gormDB
}

func parseLogLevel(level string) logger.LogLevel {
	switch level {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Warn // Default level if not provided or unrecognized
	}
}
