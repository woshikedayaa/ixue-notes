package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/woshikedayaa/ixue_note/internal/utils"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func LoginToIXue() error {
	var (
		err  error
		resp *http.Response
		//capture
		captureFailCount int
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

	for captureFailCount < config.VerifyTryCount {
		var (
			code    = make(chan string)
			errChan = make(chan error)
		)
		go captureVerify(context.Background(), code, errChan)
		// 组织数据
		form := url.Values{}
		form.Add("account", account)
		form.Add("password", password)
		form.Add("csrf_app_name", csrf)
		select {
		case captureCode := <-code:
			form.Add("verify", captureCode)
		case err = <-errChan:
			return err
		}
		// 开始登陆
		log.Printf("POST %s\n", loginUrl.Build())
		resp, err = client.POSTForm(loginUrl.Build(), form)
		if err != nil {
			return err
		}
		// 这个接口返回html和json json很小可能大于1024 所有大于 1024 我们就视为html 就直接错误
		if resp.ContentLength > 1024 {
			return errors.New("without permission to access this resource")
		}
		// 这里正常返回的json 我们绑定到model上去

		// close
		resp.Body.Close()
	}
}

func captureVerify(ctx context.Context, code chan string, errChan chan error) {
	var (
		resp *http.Response
		err  error
		// capture
		verifyUrl = utils.URLBuild("https://app.readoor.cn/app/napp/getImageCode/login/4/80/40").
				AddArg("s_id", config.AppID).
				AddArg("rand", strconv.FormatFloat(rand.Float64(), 'f', 16, 64))
		verifyCode        = ""
		verifyCodebs      []byte
		captureImageBytes []byte
	)

	// 去获取验证码
	log.Printf("GET %s\n", verifyUrl.Build())
	resp, err = client.GET(verifyUrl.Build())
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()
	captureImageBytes, err = io.ReadAll(resp.Body)
	if config.AutoVerify {
		//todo auto verify use gemini ai
		panic("auto verify 失败(work in process)")
	} else {
		log.Println("auto-verify 已经关闭,验证码图片将导出至 Capture.png")

		err = os.WriteFile("capture.png", captureImageBytes, 0666)
		if err != nil {
			errChan <- err
			return
		}
		fmt.Printf("请输入验证码:")
		_, err = fmt.Scanf("%s", &verifyCode)
		if err != nil {
			errChan <- err
			return
		}
		verifyCodebs, err = utils.RSAEncrypt([]byte(verifyCode))
		if err != nil {
			errChan <- err
			return
		}

	}
	verifyCode = base64.StdEncoding.EncodeToString(verifyCodebs)
	code <- verifyCode
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
