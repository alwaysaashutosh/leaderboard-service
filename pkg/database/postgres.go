package database

// import (
// 	"fmt"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// type postgresDB struct {
// 	*gorm.DB
// }

// // returns new postgresql db instance for provided config
// func newPostgresDB(cfg *Config) (Queries, error) {
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=UTC", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)

// 	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		PrepareStmt: true,
// 		Logger:      logger.Default.LogMode(logger.Info),
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot connect to database at %s:%v/%s due to %v", cfg.Host, cfg.Port, cfg.Name, err)
// 	}

// 	sqlDB, err := database.DB()
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot get sql db instance due to %v", err)
// 	}

// 	sqlDB.SetMaxIdleConns(cfg.IdleConn)
// 	sqlDB.SetMaxOpenConns(cfg.OpenConn)

// 	err = sqlDB.Ping()
// 	if err != nil {
// 		return nil, fmt.Errorf("pinging the database failed with %v", err)
// 	}

// 	return &postgresDB{database}, nil
// }

// // InsertData()
// //
// //	FetchRank()
// //	GetTopNRank()
// func (db *postgresDB) InsertData(details *Leaderboard) error {
// 	return db.Create(details).Error
// }

// // func (db *postgresDB) UpsertRegistryScanDetails(details *RegistryScanDetail) error {
// // 	// Retries and updates the scan Details table.
// // 	return backoff.Retry(func() error {
// // 		return db.Where(RegistryScanDetail{RegistryID: details.RegistryID, ScheduledDate: details.ScheduledDate}).FirstOrCreate(details).Error
// // 	}, backoff.WithMaxRetries(&backoff.ConstantBackOff{Interval: db.retryMaxInterval}, db.maxRetryCount))

// // }

// // func (db *postgresDB) GetRegistryScanDetails(details *RegistryScanDetail, registryScanId string) error {
// // 	return backoff.Retry(func() error {
// // 		return db.Where("id = ?", registryScanId).Find(details).Error
// // 	}, backoff.WithMaxRetries(&backoff.ConstantBackOff{Interval: db.retryMaxInterval}, db.maxRetryCount))
// // }
