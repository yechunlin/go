package lib

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RedisConnect *redis.Client
var Ctx = context.Background()

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	RedisConnect = rdb
}
