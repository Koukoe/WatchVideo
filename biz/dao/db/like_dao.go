package db

import (
	"WatchVideo/biz/model/store"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateLike(l *store.Like) (bool, error) {
	res := DB.Clauses(clause.OnConflict{DoNothing: true}).Create(l)
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected > 0, nil
}

func DeleteLike(userID, videoID string) (int64, error) {
	res := DB.Unscoped().Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&store.Like{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func GetLikeByUserAndVideo(userID, videoID string) (*store.Like, error) {
	var like store.Like
	err := DB.Where("user_id = ? AND video_id = ?", userID, videoID).First(&like).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &like, nil
}

func ListLikedVideosByUserID(userID string, offset, limit int) ([]*store.Video, error) {
	sub := DB.Model(&store.Like{}).Select("video_id").Where("user_id = ?", userID)

	var videos []*store.Video
	if err := DB.Model(&store.Video{}).
		Where("id IN (?)", sub).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func IncVideoLikeCount(videoID string) error {
	return DB.Model(&store.Video{}).
		Where("id = ?", videoID).
		UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error
}

func DecVideoLikeCount(videoID string) error {
	return DB.Model(&store.Video{}).
		Where("id = ?", videoID).
		UpdateColumn("like_count", gorm.Expr("CASE WHEN like_count > 0 THEN like_count - 1 ELSE 0 END")).Error
}

func CountLikesByUserID(userID string) (int64, error) {
	var count int64
	if err := DB.Model(&store.Like{}).
		Where("user_id = ?", userID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
