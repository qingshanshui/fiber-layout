package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	"wat.ink/layout/fiber/pkg/config"
	"wat.ink/layout/fiber/pkg/logger"
)

var (
	Client *redis.Client
	ctx    = context.Background()
)

// Init 初始化 Redis 连接
func Init() {
	if !config.Conf.Redis.Enable {
		logger.Info("Redis is disabled")
		return
	}

	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Conf.Redis.Host, config.Conf.Redis.Port),
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.Database,
	})

	// 测试连接
	if err := Client.Ping(ctx).Err(); err != nil {
		logger.Fatal("Failed to connect to Redis", "error", err)
	}

	logger.Info("Successfully connected to Redis",
		"host", config.Conf.Redis.Host,
		"port", config.Conf.Redis.Port,
	)
}

// Set 设置键值对
func Set(key string, value interface{}, expiration time.Duration) error {
	if Client == nil {
		return nil
	}
	return Client.Set(ctx, key, value, expiration).Err()
}

// Get 获取值
func Get(key string) (string, error) {
	if Client == nil {
		return "", nil
	}
	return Client.Get(ctx, key).Result()
}

// Del 删除键
func Del(keys ...string) error {
	if Client == nil {
		return nil
	}
	return Client.Del(ctx, keys...).Err()
}

// Exists 检查键是否存在
func Exists(key string) (bool, error) {
	if Client == nil {
		return false, nil
	}
	result, err := Client.Exists(ctx, key).Result()
	return result > 0, err
}

// TTL 获取键的过期时间
func TTL(key string) (time.Duration, error) {
	if Client == nil {
		return 0, nil
	}
	return Client.TTL(ctx, key).Result()
}

// Close 关闭 Redis 连接
func Close() error {
	if Client == nil {
		return nil
	}
	return Client.Close()
} 