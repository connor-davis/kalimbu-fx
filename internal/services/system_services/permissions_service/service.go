package permissions_service

import (
	"github.com/connor-davis/kalimbu-fx/internal/models/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PermissionsService struct {
	logger   *zap.Logger
	postgres *gorm.DB
}

func NewPermissionsService(logger *zap.Logger, postgres *gorm.DB) *PermissionsService {
	return &PermissionsService{
		logger:   logger,
		postgres: postgres,
	}
}

func (s *PermissionsService) Create(permission *system.SystemPermission) error {
	if err := s.postgres.
		Create(permission).Error; err != nil {
		s.logger.Error("Failed to create permission", zap.Error(err))

		return err
	}

	return nil
}

func (s *PermissionsService) Update(permission *system.SystemPermission, clauses ...clause.Expression) error {
	if err := s.postgres.
		Model(&system.SystemPermission{}).
		Clauses(clauses...).
		Updates(permission).Error; err != nil {
		s.logger.Error("Failed to update permission", zap.Error(err))

		return err
	}

	return nil
}

func (s *PermissionsService) Delete(clauses ...clause.Expression) error {
	if err := s.postgres.
		Clauses(clauses...).
		Delete(&system.SystemPermission{}).Error; err != nil {
		s.logger.Error("Failed to delete permission", zap.Error(err))

		return err
	}

	return nil
}

func (s *PermissionsService) Get(clauses ...clause.Expression) (*system.SystemPermission, error) {
	var permission *system.SystemPermission

	if err := s.postgres.
		Clauses(clauses...).
		First(&permission).Error; err != nil {
		s.logger.Error("Failed to get permission", zap.Error(err))

		return nil, err
	}

	return permission, nil
}

func (s *PermissionsService) GetAll(clauses ...clause.Expression) ([]system.SystemPermission, error) {
	var permissions []system.SystemPermission

	if err := s.postgres.
		Clauses(clauses...).
		Find(&permissions).Error; err != nil {
		s.logger.Error("Failed to get all permissions", zap.Error(err))

		return nil, err
	}

	return permissions, nil
}
