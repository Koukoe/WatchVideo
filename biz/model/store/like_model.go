package store

import (
	"time"

	"gorm.io/gorm"
)

type Like struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    string         `gorm:"type:varchar(64);uniqueIndex:idx_user_video;not null" json:"user_id"`
	VideoID   string         `gorm:"type:varchar(64);uniqueIndex:idx_user_video;not null" json:"video_id"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
