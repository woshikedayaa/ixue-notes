package main

import (
	"fmt"
	"github.com/woshikedayaa/ixue_note/internal/utils"
	"log"
	"net/http"
)

func LoginToIXue() error {
	// init
	err := initAccessIXue()
	if err != nil {
		return err
	}
}

func initAccessIXue() error {
	u := utils.URLBuild("https://app.readoor.cn/app/dt/pd/1544059443").
		AddArg("s", "1").
		Build()
	log.Printf("GET %s\n", u)
	resp, err := client.GET(u)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("LoginToIXue: status code %d", resp.StatusCode)
	}

	log.Println("获取到了 csrf-cookie: ", client.GetCookie())

	return nil
}
