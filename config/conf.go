package conf

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func init() {
	// conf := Conf{}
	// conf.InitConfigYaml()
	workDir, _ := os.Getwd()
	fmt.Println(workDir, "work")
	viper.SetConfigName("config.dev")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("读取失败-------------------------------")
}
