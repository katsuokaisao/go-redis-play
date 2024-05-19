package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/katsuokaisao/go-redis-play/domain"
)

func ProvideLocalRedis(cfg *domain.RedisCfg) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})
}

func ProvideRemoteRedis(cfg *domain.RedisCfg) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Username: cfg.Username,
		Password: cfg.Password,
	})
}
