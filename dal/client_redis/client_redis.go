package client_redis

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/go-redis/redis/v8"

	"group_shopping_mall/utils/utils"
)

var redisClient *redis.Client = nil

type RedisConfig struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Addr     string `json:"addr,omitempty"`
	Database int    `json:"database,omitempty"`
}

func InitRedis() {
	config := utils.GetConfig[RedisConfig]("redisConfig.json")
	ctx := context.Background()
	redisClient = redis.NewClient(&redis.Options{
		Username: config.Username,
		Addr:     config.Addr,
		Password: config.Password, // 没有密码，默认值
		DB:       config.Database, // 默认DB 0
	})

	// 下述代码为您提供SET与GET的使用示例。
	err := redisClient.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("set : foo -> ", val)
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		InitRedis()
	}
	return redisClient
}

func SetValue(ctx context.Context, key string, value any, duration time.Duration) error {
	client := GetRedisClient()
	err := client.Set(ctx, key, value, duration).Err()
	return err
}

func GetValue(ctx context.Context, key string) (string, error) {
	client := GetRedisClient()
	val, err := client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", errors.Errorf("get value from redis failed, key: %s, err: %v", key, err)
	}

	return val, nil
}
