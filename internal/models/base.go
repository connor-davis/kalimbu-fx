package models

import "github.com/google/uuid"

type Base struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt string    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt string    `gorm:"autoUpdateTime" json:"updated_at"`
}
