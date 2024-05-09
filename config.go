package main

import (
	_ "embed"
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
)

//go:embed config_full.yaml
var configExampleFUll string

var maq = "_uid,,_appId,1544059443,_cid,89,_isMobile,,_isWeixin,9.1,_accessType,,_keyValue,,_appver,2.29.1,_apiver,1.0"

type Config struct {
	User   user   `yaml:"user"`
	Target target `yaml:"target"`

	// why mapstructure
	// see : https://gist.github.com/chazcheadle/45bf85b793dea2b71bd05ebaa3c28644
	HttpTimeOut int    `mapstructure:"http-timeout"`
	BaseUrl     string `mapstructure:"base-url"`
	AppID       string `mapstructure:"app-id"`
}

type user struct {
	Account   string `yaml:"account"`
	Password  string `yaml:"password"`
	Encoded   bool   `yaml:"encoded"`
	Csrf      string `yaml:"csrf"`
	SessionID string `yaml:"session-id"`
}

type target struct {
	Time   int `yaml:"time"`
	Finish int `yaml:"finish"`
	BookID int `yaml:"book_id"`
}

var config Config

func ConfigInit() error {

	// 第一次的时候创建文件
	var err error
	var stat os.FileInfo
	stat, err = os.Stat("config.yaml")
	if os.IsNotExist(err) {
		err = nil
		err = os.WriteFile("config.yaml", []byte(configExampleFUll), 0644)
		if err != nil {
			log.Println(err)
			_ = os.Remove("config.yaml")
		} else {
			log.Println("file config.yaml created")
		}
		os.Exit(0)
	}
	// 检查问题
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return errors.New("config.yaml is a directory")
	}
	if stat.Size() == 0 {
		return errors.New("config.yaml is empty")
	}
	// 最后开始正式读配置文件
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return viper.Unmarshal(&config)
}
