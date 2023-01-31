package initalize

import (
	"context"
	"fmt"
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
	var ctx = context.Background()
	_, err := Rdb.Ping(ctx).Result()

	if err != nil {
		panic("Redis 链接 失败：" + err.Error())
	}
	fmt.Println("redis 链接成功")
}
