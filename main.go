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
	log.Println("获取到了 csrf_app_name:", config.User.Csrf)
	log.Println("获取到了 vpapp_session:", config.User.SessionID)
	// 这里请求完毕了再添加maq
	log.Println("_maq 将默认使用", maq)
	client.AddCookie([]*http.Cookie{{
		Name:    "_maq",
		Value:   maq,
		Path:    "/",
		Domain:  config.BaseUrl,
		MaxAge:  math.MaxInt,
		Expires: time.Now().Add(time.Hour * 24 * 365), // 一年日期 不信它会过期(
	},
	})
	// 登陆
	err = LoginToIXue()
	if err != nil {
		panic(err)
	}
	//-----------------
	// 这里本来是想让用户选一个页面来刷的 然后 发现 太麻烦了 主要是解析返回的json不方便
	// 然后后面还有一大堆东西 就先摆烂了。。。。。
	//-----------------
	//// 这里开始获取书的列表
	//err = GetBooKList()
	//if err != nil {
	//	panic(err)
	//}
	//log.Printf("准备刷的书是 %s 准备刷的时间是 %dmin\n", TargetBookInfo.Data.BookName, config.Target.Time)
	// 3：55
	train(60, 10)
	log.Println("学习完毕!")
	select {}
}
