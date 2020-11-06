package config

import (
	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
)

type Config struct {
}

var ConfigInstance = Config{}

func init() {
	// init ConfigInstance
}

func (c *Config) GetRedisOptions() *redis.Options {
	return &redis.Options{}
}

func (c *Config) GetMysqlConfig() *mysql.Config {
	return &mysql.Config{}
}
