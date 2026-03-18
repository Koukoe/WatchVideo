package db

import (
	"WatchVideo/biz/model/store"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	d, err := gorm.Open(sqlite.Open("watchvideo.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := d.AutoMigrate(&store.User{}); err != nil {
		return err
	}
	if err := d.AutoMigrate(&store.Video{}); err != nil {
		return err
	}
	DB = d
	return nil
}
