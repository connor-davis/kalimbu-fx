package system

import "github.com/connor-davis/kalimbu-fx/internal/models"

type SystemRole struct {
	models.Base
	Name        string             `gorm:"type:text;not null;uniqueIndex;" json:"name"`
	Description string             `gorm:"type:text;" json:"description,omitempty"`
	IsDefault   bool               `gorm:"default:false" json:"isDefault"`
	Permissions []SystemPermission `gorm:"foreignKey:RoleId;references:Id;constraint:OnDelete:CASCADE;" json:"permissions,omitempty"`
}
