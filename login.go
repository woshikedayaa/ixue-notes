package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/woshikedayaa/ixue_note/internal/models"
	"github.com/woshikedayaa/ixue_note/internal/utils"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var userDetailedInfo *models.UserDetailedInfo

func LoginToIXue() error {
	var (
		err  error
		resp *http.Response
		// account
		loginUrl = utils.URLBuild("https://app.readoor.cn/app/cm/login").
				AddArg("s_id", config.AppID)
		account  = config.User.Account
		password = config.User.Password
		code     = ""
	)
	account, password, code, err = getUserInfo()
	if err != nil {
		return err
	}
	// 组织数据
	form := url.Values{}
	form.Add("account", account)
	form.Add("password", password)
	form.Add("csrf_app_name", config.User.Csrf)
	form.Add("verify", code)
	// 开始登陆
	log.Printf("POST %s\n", loginUrl.Build())
	resp, err = client.POSTForm(loginUrl.Build(), form)
	if err != nil {
		return err
	}

	// close
	defer resp.Body.Close()

	var j []byte
	j, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// 这个接口返回html和json json很小可能大于1024 所有大于 1024 我们就视为html 就直接错误
	if utils.IsHtml(int(resp.ContentLength)) {
		return errors.New("without permission to access this resource")
	}

	// 这里正常返回的json 我们绑定到model上去
	return json.Unmarshal(j, &userDetailedInfo)
}

func getUserInfo() (account string, password string, captureCode string, err error) {

	if len(config.User.Account) == 0 || len(config.User.Password) == 0 {
		if len(config.User.Account) == 0 {
			fmt.Printf("请输入账号:")
			_, err = fmt.Scanf("%s", &account)
			if err != nil {
				return "", "", "", err
			}
		}
		fmt.Printf("请输入密码:")
		_, err = fmt.Scanf("%s", &password)
		if err != nil {
			return "", "", "", err
		}
		config.User.Encoded = false
	} else {
		account = config.User.Account
		password = config.User.Password
	}
	var (
		code    = make(chan string)
		errChan = make(chan error)
	)
	// 这里把这个验证码请求放到输入账密后 防止输出乱套
	go captureVerify(context.Background(), code, errChan)

	// 加密数据
	if !config.User.Encoded {
		account, err = utils.RSAEncrypt([]byte(account))
		if err != nil {
			return "", "", "", err
		}
		password, err = utils.RSAEncrypt([]byte(password))
		if err != nil {
			return "", "", "", err
		}
	}

	select {
	case code := <-code:
		captureCode = code
	case err := <-errChan:
		return "", "", "", err
	}
	return account, password, captureCode, nil
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
	// 手动识别验证码
	captureImageBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		errChan <- err
		return
	}
	err = os.WriteFile("capture.png", captureImageBytes, 0666)
	if err != nil {
		errChan <- err
		return
	}
	log.Println("验证码已经导出至: capture.png")
	fmt.Printf("请输入验证码:")
	_, err = fmt.Scanf("%s", &verifyCode)
	if err != nil {
		errChan <- err
		return
	}
	verifyCode, err = utils.RSAEncrypt([]byte(verifyCode))
	if err != nil {
		errChan <- err
		return
	}
	// 返回数据
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
	config.User.Csrf = client.GetCookie(u.BuildUrl()).Find("csrf_cookie_name").Value
	config.User.SessionID = client.GetCookie(u.BuildUrl()).Find("vpapp_session").Value
	return nil
}
