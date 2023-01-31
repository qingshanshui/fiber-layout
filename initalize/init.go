package initalize

import (
	"github.com/spf13/viper"
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		if viper.GetBool("Mysql.Enable") {
			InitDatabaseMysql()
		}
		if viper.GetBool("Redis.Enable") {
			initDatabaseRedis()
		}
	})
}
