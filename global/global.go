package global

import (
	"github.com/redis/go-redis/v9"
)

var (
	// Redis Redis连接
	Redis *redis.Client
)
