package db

import (
	"errors"

	"WatchVideo/biz/model/store"

	"gorm.io/gorm"
)

// 创建用户
func CreateUser(u *store.User) error {
	return DB.Create(u).Error
}

// 根据用户名获取用户信息
func GetUserByUsername(username string) (*store.User, error) {
	var u store.User
	err := DB.Where("username = ?", username).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// 根据用户ID获取用户信息
func GetUserByID(id string) (*store.User, error) {
	var u store.User
	err := DB.Where("id = ?", id).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// 更新用户头像URL
func UpdateUserAvatar(userID, avatarURL string) error {
	return DB.Model(&store.User{}).Where("id = ?", userID).
		Update("avatar_url", avatarURL).Error
}

// 获取下一个用户ID
func NextUserID() (int64, error) {
	var count int64
	if err := DB.Model(&store.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return 10001 + count, nil
}
