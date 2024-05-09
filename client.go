package main

import (
	"github.com/woshikedayaa/ixue_note/internal"
	"net/http"
	"time"
)

var client *internal.HttpClientWrapper

func HttpClientInit() {
	hpc := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
		Timeout: time.Duration(config.HttpTimeOut) * time.Second,
	}

	client = internal.NewHttpClientWrapper(&hpc)
	client.EnableRandomUA()
	client.EnableRandomReferer()
}
