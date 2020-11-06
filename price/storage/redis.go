package storage

import (
	"github.com/go-redis/redis"
	"github.com/gtforge/isavinof/test/config"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(config.ConfigInstance.GetRedisOptions())
}
