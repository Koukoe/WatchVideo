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
