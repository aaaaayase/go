package config

import (
	"github.com/go-redis/redis"
	"log"
	"week1/global"
)

func initRedis() {
	RedisClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    AppConfig.Redis.Addrs,
		Password: AppConfig.Redis.Password,
		PoolSize: AppConfig.Redis.PoolSize,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("redis init err: %v", err)
	}

	global.RedisDB = RedisClient
}
