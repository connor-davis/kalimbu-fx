package postgres

import (
	"github.com/connor-davis/kalimbu-fx/env"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase(logger *zap.Logger) *gorm.DB {
	logger.Info("Connecting to the database...")

	db, err := gorm.Open(postgres.Open(string(env.POSTGRES_DSN)), &gorm.Config{})

	if err != nil {
		logger.Error("Failed to connect to the database", zap.Error(err))

		return nil
	}

	logger.Info("Connected to the database successfully")

	return db
}
