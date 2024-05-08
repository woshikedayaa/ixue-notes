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

type Config struct {
	User   user   `yaml:"user"`
	Target target `yaml:"target"`

	AutoVerify  bool `yaml:"auto-verify"`
	HttpTimeOut int  `yaml:"http-timeout"`
}

type user struct {
	account  string `yaml:"account"`
	password string `yaml:"password"`
	encoded  bool   `yaml:"encoded"`
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
	viper.SetConfigFile("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return viper.Unmarshal(&config)
}
