package migrate

import (
	"github.com/connor-davis/kalimbu-fx/internal/models/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Migrations struct{}

func RunMigrations(logger *zap.Logger, postgres *gorm.DB) *Migrations {
	logger.Info("Running database migrations...")

	if err := postgres.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		logger.Error("Failed to create uuid-ossp extension", zap.Error(err))

		return nil
	}

	postgres.AutoMigrate(
		&system.SystemRole{},
		&system.SystemUser{},
		&system.SystemPermission{},
	)

	logger.Info("Database migrations completed successfully")

	return &Migrations{}
}
