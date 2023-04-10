package initalize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var Rdb *redis.Client

func initDatabaseRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("Redis.Host") + ":" + viper.GetString("Redis.Port"),
		Password: viper.GetString("Redis.Password"),
		DB:       viper.GetInt("Redis.Database"),
	})
	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("Redis 链接 失败：" + err.Error())
	}
}
