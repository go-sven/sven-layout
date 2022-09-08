package data

import (
	"errors"
	"github/go-sven/sven-layout/app/conf"
	"github.com/go-redis/redis"
	"time"
)

func NewRdb(config *conf.RedisConfig) *redis.Client  {
	addr := config.Host+ ":" +config.Port
	redisClient := redis.NewClient(&redis.Options{
		Addr:    addr, // Redis地址
		Password: config.Password,  // Redis账号
		DB:       config.Db,   // Redis库
		PoolSize: 10,  // Redis连接池大小
		MaxRetries: 3,              // 最大重试次数
		IdleTimeout: 10*time.Second,            // 空闲链接超时时间
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	if err == redis.Nil {
		panic(errors.New("Redis abnormal "))
	}
	return redisClient
}