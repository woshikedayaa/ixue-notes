package main

import (
	"log"
	"math"
	"net/http"
	"time"
)

func init() {
	err := ConfigInit()
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	HttpClientInit()
	// init
	err = InitAccessIXue()
	if err != nil {
		panic(err)
	}

	log.Println("获取到了 csrf_app_name:", csrf)
	log.Println("获取到了 vpapp_session:", vpappSession)
	// 这里请求完毕了再添加maq
	maq := "_uid,,_appId,1544059443,_cid,89,_isMobile,,_isWeixin,9.1,_accessType,,_keyValue,,_appver,2.29.1,_apiver,1.0"
	log.Println("_maq 将默认使用", maq)
	client.AddCookie([]*http.Cookie{
		&http.Cookie{
			Name:    "_maq",
			Value:   maq,
			Path:    "/",
			Domain:  config.BaseUrl,
			MaxAge:  math.MaxInt,
			Expires: time.Now().Add(time.Hour * 24 * 365), // 一年日期 不信它会过期(
		},
	})
	err = LoginToIXue()
}
