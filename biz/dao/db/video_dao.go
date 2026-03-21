package db

import (
	"WatchVideo/biz/model/store"
)

func CreateVideo(v *store.Video) error {
	return DB.Create(v).Error
}

func NextVideoID() (int64, error) {
	var count int64
	if err := DB.Model(&store.Video{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return 10001 + count, nil
}

func ListVideosByAuthorID(authorID string, offset, limit int) ([]*store.Video, error) {
	var videos []*store.Video
	if err := DB.Where("author_id = ?", authorID).Offset(offset).Limit(limit).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func CountVideosByAuthorID(authorID string) (int64, error) {
	var count int64
	if err := DB.Model(&store.Video{}).Where("author_id = ?", authorID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
