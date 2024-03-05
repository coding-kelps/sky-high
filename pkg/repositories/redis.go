package repositories

import (
	"github.com/redis/go-redis/v9"

	"github.com/coding-kelps/sky-high/pkg/config"
)

var REDIS *redis.Client

func InitRedisClient(cfg *config.RedisConfig) {
	REDIS = redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
		Password: cfg.Password,
		DB:   cfg.DB,
	})
}
