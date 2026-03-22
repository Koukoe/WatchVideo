package cache

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
	Ctx = context.Background()
)

func Init() error {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}

	password := os.Getenv("REDIS_PASSWORD")

	db := 0
	if dbStr := os.Getenv("REDIS_DB"); dbStr != "" {
		if v, err := strconv.Atoi(dbStr); err == nil {
			db = v
		}
	}

	RDB = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		DialTimeout:  3 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	})

	return RDB.Ping(Ctx).Err()
}

func Close() error {
	if RDB == nil {
		return nil
	}
	return RDB.Close()
}
