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

func ListVideosByUserID(userID string, offset, limit int) ([]*store.Video, error) {
	var videos []*store.Video
	if err := DB.Where("author_id = ?", userID).Offset(offset).Limit(limit).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func CountVideosByUserID(userID string) (int64, error) {
	var count int64
	if err := DB.Model(&store.Video{}).Where("author_id = ?", userID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func SearchVideos(keywords, username, fromDate, toDate string, offset, limit int) ([]*store.Video, error) {
	query := DB.Model(&store.Video{})

	if keywords != "" {
		likePattern := "%" + keywords + "%"
		query = query.Where("title LIKE ? OR description LIKE ?", likePattern, likePattern)
	}

	if username != "" {
		likePattern := "%" + username + "%"
		query = query.Where("author_id IN (SELECT id FROM users WHERE username LIKE ?)", likePattern)
	}

	if fromDate != "" {
		query = query.Where("created_at >= ?", fromDate)
	}
	if toDate != "" {
		query = query.Where("created_at <= ?", toDate)
	}

	var videos []*store.Video
	if err := query.Offset(offset).Limit(limit).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func CountSearchVideos(keywords, username, fromDate, toDate string) (int64, error) {
	query := DB.Model(&store.Video{})

	if keywords != "" {
		likePattern := "%" + keywords + "%"
		query = query.Where("title LIKE ? OR description LIKE ?", likePattern, likePattern)
	}

	if username != "" {
		likePattern := "%" + username + "%"
		query = query.Where("author_id IN (SELECT id FROM users WHERE username = ?)", likePattern)
	}

	if fromDate != "" {
		query = query.Where("created_at >= ?", fromDate)
	}
	if toDate != "" {
		query = query.Where("created_at <= ?", toDate)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
