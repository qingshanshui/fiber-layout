package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func init() {
	var mode string
	flag.StringVar(&mode, "mode", "dev", "开发dev，生产prod")
	flag.Parse()
	workDir, _ := os.Getwd()
	viper.SetConfigName("config." + mode)
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变化的文件：", e.Name)
	})
}
