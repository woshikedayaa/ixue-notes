package main

import (
	"encoding/base64"
	"fmt"
	"github.com/woshikedayaa/ixue_note/internal/utils"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
)

func LoginToIXue() error {
	var (
		err  error
		resp *http.Response
		// capture
		verifyUrl = utils.URLBuild("https://app.readoor.cn/app/napp/getImageCode/login/4/80/40").
				AddArg("s_id", config.AppID).
				AddArg("rand", strconv.FormatFloat(rand.Float64(), 'f', 16, 64))
		verifyCode        = ""
		captureImageBytes []byte
		verifyFailCount   int

		// account
		loginUrl = utils.URLBuild("https://app.readoor.cn/app/cm/login").
				AddArg("s_id", config.AppID)
		account  = config.User.account
		password = config.User.password
	)

	if len(account) == 0 || len(password) == 0 {
		if len(account) == 0 {
			fmt.Printf("请输入账号:")
			_, err = fmt.Scanf("%s", &account)
			if err != nil {
				return err
			}
		}
		fmt.Println("请输入密码:")
		_, err = fmt.Scanf("%s", &password)
		if err != nil {
			return err
		}
		config.User.encoded = false
	}
	if !config.User.encoded {
		accountbs, err := utils.RSAEncrypt([]byte(account))
		if err != nil {
			return err
		}
		passwordbs, err := utils.RSAEncrypt([]byte(password))
		if err != nil {
			return err
		}
		account = base64.StdEncoding.EncodeToString(accountbs)
		password = base64.StdEncoding.EncodeToString(passwordbs)
	}

	for verifyFailCount <= config.VerifyTryCount {
		// 去获取验证码
		log.Printf("GET %s\n", verifyUrl.Build())
		resp, err = client.GET(verifyUrl.Build())
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		captureImageBytes, err = io.ReadAll(resp.Body)
		if config.AutoVerify {
			//todo auto verify
			panic("auto verify 失败(work in process)")
		} else {
			log.Println("auto-verify 已经关闭,验证码图片将导出至 Capture.png")

			err = os.WriteFile("capture.png", captureImageBytes, 0666)
			if err != nil {
				return err
			}
			fmt.Printf("请输入验证码:")
			_, err = fmt.Scanf("%s", &verifyCode)
			if err != nil {
				return err
			}
			verifyCodebs, err := utils.RSAEncrypt([]byte(verifyCode))
			if err != nil {
				return err
			}
			verifyCode = base64.StdEncoding.EncodeToString(verifyCodebs)
		}
		// 开始登陆
	}
}
func verifyCapture() (code string, err error) {

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
