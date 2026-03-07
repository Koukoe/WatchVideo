package store

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string         `gorm:"primaryKey;type:varchar(64)" json:"id"`
	Username     string         `gorm:"uniqueIndex;type:varchar(64);not null" json:"username"`
	PasswordHash string         `gorm:"type:varchar(255);not null" json:"-"`
	AvatarURL    string         `gorm:"type:text" json:"avatar_url"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
