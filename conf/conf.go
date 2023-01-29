package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func init() {
	conf := Conf{}
	conf.InitConfigYaml()
}

type Conf struct {
	App struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Os struct {
		DoMain string `yaml:"domain"`
	}
	Mysql struct {
		Enable   bool   `yaml:"enable"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
	}
	Redis struct {
		Enable   bool   `yaml:"enable"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database int    `yaml:"database"`
		Password string `yaml:"password"`
	}
	Zap struct {
		Filename   string `yaml:"filename"`
		MaxSize    int    `yaml:"maxSize"`
		MaxBackups int    `yaml:"maxBackups"`
		MaxAge     int    `yaml:"maxAge"`
		Compress   bool   `yaml:"compress"`
	}
	Jwt struct {
		Secret string `yaml:"secret"`
		Expire int64  `yaml:"expire"`
	}
	Md5 struct {
		Hash string `yaml:"hash"`
	}
	Debug bool `yaml:"debug"`
}

// Config 全局config
var Config = &Conf{}

func (t *Conf) InitConfigYaml() {
	// 读取配置文件
	content, err := ioutil.ReadFile("./config.dev.yaml") // 本地配置
	//content, err := ioutil.ReadFile("./config.pord.yaml") // 线上配置
	if err != nil {
		panic("读取 config.yaml 读取错误")
	}
	// yaml.Unmarshal 解析配置 到结构体
	if yaml.Unmarshal(content, t) != nil {
		panic("解析config.yaml出错")
	} else {
		fmt.Println("读取成功", t)
	}
	Config = t
}
