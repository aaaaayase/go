package config

import (
	"github.com/go-redis/redis"
	"log"
	"week1/global"
)

func initRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Addr,
		Password: AppConfig.Redis.Password,
		DB:       AppConfig.Redis.DB,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("redis init err: %v", err)
	}

	global.RedisDB = RedisClient
}
