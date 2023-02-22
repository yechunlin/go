package model

import (
	"api/conf"
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisConnect *redis.Client
var Ctx = context.Background()

func RedisInit() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.REDIS_ADDR,
		Password: conf.REDIS_PASSWORD, // no password set
		DB:       conf.REDIS_DB,       // use default DB
	})
	RedisConnect = rdb
}
