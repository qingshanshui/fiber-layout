package initalize

import (
	"fiber-layout/conf"
	"fiber-layout/pkg/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabaseMysql() {
	//初始化数据库
	mysqlUrl := utils.StringBytesBufferJoin(conf.Config.Mysql.User, ":", conf.Config.Mysql.Password, "@tcp(", conf.Config.Mysql.Host, ":",
		conf.Config.Mysql.Port, ")/", conf.Config.Mysql.DbName, "?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlUrl, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("mysql数据库链接失败gorm")
	}
	fmt.Println("mysql 链接成功")
	DB = db
}
