package utils

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var (
	Redis_Context = context.Background()
	DB_Redis *redis.Client
)

func init() {
	// 初始化 Redis 客户端
	DB_Redis = redis.NewClient(&redis.Options{
		Addr:     "172.25.13.23:6379", // Redis 地址
		Password: "123",               // Redis 密码 (如果没有可以为空)
		DB:       0,                    // 默认 DB
	})

	// 测试 Ping 命令，确保连接正常
	_, err := DB_Redis.Ping(Redis_Context).Result()

	if err == nil {
		fmt.Println("Redis 连接成功")
	} else {
		fmt.Printf("Redis 连接失败: %v\n", err)
	}
}
