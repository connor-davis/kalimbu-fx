package usersService

import (
	"github.com/connor-davis/kalimbu-fx/internal/models/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UsersService struct {
	logger   *zap.Logger
	postgres *gorm.DB
}

func NewUsersService(logger *zap.Logger, postgres *gorm.DB) *UsersService {
	return &UsersService{
		logger:   logger,
		postgres: postgres,
	}
}

func (s *UsersService) Create(user *system.SystemUser) error {
	if err := s.postgres.
		Create(user).Error; err != nil {
		s.logger.Error("Failed to create user", zap.Error(err))

		return err
	}

	return nil
}

func (s *UsersService) Update(user *system.SystemUser, clauses ...clause.Expression) error {
	if err := s.postgres.
		Model(&system.SystemUser{}).
		Clauses(clauses...).
		Updates(user).Error; err != nil {
		s.logger.Error("Failed to update user", zap.Error(err))

		return err
	}

	return nil
}

func (s *UsersService) Delete(clauses ...clause.Expression) error {
	if err := s.postgres.
		Clauses(clauses...).
		Delete(&system.SystemUser{}).Error; err != nil {
		s.logger.Error("Failed to delete user", zap.Error(err))

		return err
	}

	return nil
}

func (s *UsersService) Get(clauses ...clause.Expression) (*system.SystemUser, error) {
	var user *system.SystemUser

	if err := s.postgres.
		Clauses(clauses...).
		Find(&user).Error; err != nil {
		s.logger.Error("Failed to get user", zap.Error(err))

		return nil, err
	}

	return user, nil
}

func (s *UsersService) GetAll(clauses ...clause.Expression) ([]system.SystemUser, error) {
	var users []system.SystemUser

	if err := s.postgres.
		Clauses(clauses...).
		Find(&users).Error; err != nil {
		s.logger.Error("Failed to get all users", zap.Error(err))

		return nil, err
	}

	return users, nil
}
