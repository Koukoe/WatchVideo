package store

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	ID           string         `gorm:"primaryKey;type:varchar(64)" json:"id"`
	AuthorID     string         `gorm:"type:varchar(64);index;not null" json:"author_id"`
	Title        string         `gorm:"type:varchar(128);not null" json:"title"`
	Description  string         `gorm:"type:text" json:"description"`
	PlayURL      string         `gorm:"type:text;not null" json:"play_url"`
	CoverURL     string         `gorm:"type:text" json:"cover_url"`
	LikeCount    int64          `gorm:"default:0" json:"like_count"`
	CommentCount int64          `gorm:"default:0" json:"comment_count"`
	VisitCount   int64          `gorm:"default:0" json:"visit_count"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
