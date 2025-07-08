package system

import (
	"github.com/connor-davis/kalimbu-fx/internal/models"
	"github.com/google/uuid"
)

type SystemUser struct {
	models.Base
	Email       string             `gorm:"type:text;not null;uniqueIndex;" json:"email"`
	MfaSecret   string             `gorm:"type:text;" json:"mfaSecret,omitempty"`
	MfaEnabled  bool               `gorm:"default:false" json:"mfaEnabled"`
	MfaVerified bool               `gorm:"default:false" json:"mfaVerified"`
	RoleId      uuid.UUID          `gorm:"type:uuid;" json:"-"`
	Role        SystemRole         `gorm:"foreignKey:RoleId;references:Id;constraint;onDelete:CASCADE;" json:"role,omitzero"`
	Permissions []SystemPermission `gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE;" json:"permissions,omitempty"`
}
