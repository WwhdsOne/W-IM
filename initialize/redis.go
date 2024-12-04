package initialize

import (
	"W-IM/config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func InitRedisClient(cfg *config.Config) *redis.Client {
	r := cfg.Redis
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", r.Addr, r.Port),
		Password:     r.Password,
		DB:           r.DB,
		PoolSize:     r.PoolSize,
		MinIdleConns: r.MinIdleConns,
	})
	return client
}
