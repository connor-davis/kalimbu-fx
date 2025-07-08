package system

import "github.com/connor-davis/kalimbu-fx/internal/models"

type SystemPermission struct {
	models.Base
	UserId      string `gorm:"type:text;" json:"userId,omitempty"`
	RoleId      string `gorm:"type:text;" json:"roleId,omitempty"`
	Value       string `gorm:"type:text;not null;uniqueIndex;" json:"value"`
	Description string `gorm:"type:text;not null;" json:"description"`
}
