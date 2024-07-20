package database

import (
	"TaskManagerGoLang/models"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

/*
func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	return rdb
}*/

func NewRedisClient(cfg models.Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
