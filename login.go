package main

import (
	"fmt"
	"github.com/woshikedayaa/ixue_note/internal/utils"
	"log"
	"net/http"
)

func LoginToIXue() error {
	u := utils.URLBuild("https://app.readoor.cn/app/cm/login")
}

func InitAccessIXue() error {
	u := utils.URLBuild("https://app.readoor.cn/app/dt/pd/1544059443").
		AddArg("s", "1")
	log.Printf("GET %s\n", u.Build())
	resp, err := client.GET(u.Build())
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("LoginToIXue: status code %d", resp.StatusCode)
	}
	csrf = client.GetCookie(u.BuildUrl()).Find("csrf_cookie_name").Value
	vpappSession = client.GetCookie(u.BuildUrl()).Find("vpapp_session").Value
	return nil
}
