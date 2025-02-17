package config

import (
	"flag"
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	config *Config
	once   sync.Once
)

// Config 配置结构
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	MySQL    MySQLConfig    `mapstructure:"mysql"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Log      LogConfig      `mapstructure:"log"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	MD5      MD5Config      `mapstructure:"md5"`
	RabbitMQ RabbitMQConfig `mapstructure:"rabbitmq"`
	Email    EmailConfig    `mapstructure:"email"`
	Debug    bool           `mapstructure:"debug"`
}

type AppConfig struct {
	Name string
	Mode string
	Port int
}

type MySQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Enable   bool
	Host     string
	Port     int
	Database int
	Password string
}

type LogConfig struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type JWTConfig struct {
	Secret string
	Expire int
}

type MD5Config struct {
	Hash string
}

type RabbitMQConfig struct {
	Enable   bool
	Username string
	Password string
	Host     string
	Port     int
}

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

var Conf = new(Config)

func Init() {
	var mode string
	flag.StringVar(&mode, "mode", "dev", "运行模式：dev开发环境，prod生产环境")
	flag.Parse()

	viper.SetConfigName(fmt.Sprintf("config.%s", mode))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}

	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("解析配置文件失败: %s", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("配置文件重载失败: %s\n", err)
		}
	})
}

// GetString 获取字符串配置
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt 获取整数配置
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetBool 获取布尔配置
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetConfig 获取配置单例
func GetConfig() *Config {
	return Conf
}

// 热重载配置
func WatchConfig(callback func(e fsnotify.Event)) {
	viper.WatchConfig()
	viper.OnConfigChange(callback)
}
