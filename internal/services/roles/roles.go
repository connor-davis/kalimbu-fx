package roles

import (
	"github.com/connor-davis/kalimbu-fx/internal/models/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RolesService struct {
	logger   *zap.Logger
	postgres *gorm.DB
}

func NewRolesService(logger *zap.Logger, postgres *gorm.DB) *RolesService {
	return &RolesService{
		logger:   logger,
		postgres: postgres,
	}
}

func (s *RolesService) Create(role *system.SystemRole) error {
	if err := s.postgres.
		Create(role).Error; err != nil {
		s.logger.Error("Failed to create role", zap.Error(err))

		return err
	}

	return nil
}

func (s *RolesService) Update(role *system.SystemRole, clauses ...clause.Expression) error {
	if err := s.postgres.
		Model(&system.SystemRole{}).
		Clauses(clauses...).
		Updates(role).Error; err != nil {
		s.logger.Error("Failed to update role", zap.Error(err))

		return err
	}

	return nil
}

func (s *RolesService) Delete(clauses ...clause.Expression) error {
	if err := s.postgres.
		Clauses(clauses...).
		Delete(&system.SystemRole{}).Error; err != nil {
		s.logger.Error("Failed to delete role", zap.Error(err))

		return err
	}

	return nil
}

func (s *RolesService) Get(clauses ...clause.Expression) (*system.SystemRole, error) {
	var role *system.SystemRole

	if err := s.postgres.
		Clauses(clauses...).
		Find(&role).Error; err != nil {
		s.logger.Error("Failed to get role", zap.Error(err))

		return nil, err
	}

	return role, nil
}

func (s *RolesService) GetAll(clauses ...clause.Expression) ([]system.SystemRole, error) {
	var roles []system.SystemRole

	if err := s.postgres.
		Clauses(clauses...).
		Find(&roles).Error; err != nil {
		s.logger.Error("Failed to get roles", zap.Error(err))

		return nil, err
	}

	return roles, nil
}
