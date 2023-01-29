package initalize

import (
	"fiber-layout/conf"
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		if conf.Config.Mysql.Enable {
			InitDatabaseMysql()
		}
		if conf.Config.Redis.Enable {
			initDatabaseRedis()
		}
	})
}
