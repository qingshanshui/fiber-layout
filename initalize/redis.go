package initalize

import (
	"context"
	"fiber-layout/conf"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func initDatabaseRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Config.Redis.Host + ":" + conf.Config.Redis.Port,
		Password: conf.Config.Redis.Password,
		DB:       conf.Config.Redis.Database,
	})
	var ctx = context.Background()
	_, err := Rdb.Ping(ctx).Result()

	if err != nil {
		panic("Redis 链接 失败：" + err.Error())
	}
	fmt.Println("redis 链接成功")
}
