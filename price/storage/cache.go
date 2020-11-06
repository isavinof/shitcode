package storage

import (
	"encoding/json"
	"time"
)

type Cache struct {
}

var CacheInstance = &Cache{}

func (c *Cache) Set(symbol string, price *Price) error {
	return RedisClient.Set("cache-symbol-"+symbol, price, 2*time.Hour).Err()
}

func (c *Cache) Get(symbol string) (*Price, error) {
	res, err := RedisClient.Get("cache-symbol-" + symbol).Result()
	if err != nil {
		return &Price{}, err
	}
	price := &Price{}
	err = json.Unmarshal([]byte(res), price)
	return price, err
}
