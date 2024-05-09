package main

import (
	"io"
	"log"
	"math/rand/v2"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func train(t int, count int) {
	var (
		// 这里放做题那个页面
		getInfoUrl = ""
		// ${unique_idf}
		// ${time}
		// ${submit_unique_num}
		submitData           = ``
		system               = ""
		cid                  = ""
		bookid               = ""
		previousMaterialList = ``
		materialRecordList   = ``
		setRandomCookieUrl   = ""
	)

	for i := 0; i < count; i++ {
		// 设置它自带的一个随机的cookie
		log.Printf("GET %s\n", setRandomCookieUrl)
		_, err := client.GET(setRandomCookieUrl)
		if err != nil {
			log.Println("Error:", err)
			continue
		}
		// 获取信息
		log.Printf("GET %s\n", getInfoUrl)
		resp, err := client.GET(getInfoUrl)
		if err != nil {
			log.Println("Error: ", err)
			continue
		}

		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error: ", err)
			continue
		}
		// 提取
		var (
			referrerReg  = regexp.MustCompile("var referrer_url = (.*)")
			begTimeReg   = regexp.MustCompile("var MOOCBeginTime = (.*)")
			uniqueIdfReg = regexp.MustCompile("var uidffqbsub = (.*)")
			begTime      = begTimeReg.FindString(string(bytes))
			uniqueIdf    = uniqueIdfReg.FindString(string(bytes))
			referrer     = referrerReg.FindString(string(bytes))
		)
		resp.Body.Close()
		begTime = strings.TrimSpace(begTime)
		uniqueIdf = strings.TrimSpace(uniqueIdf)
		referrer = strings.TrimSpace(referrer)

		begTime = begTime[strings.Index(begTime, "'")+1 : len(begTime)-2]
		uniqueIdf = uniqueIdf[strings.Index(uniqueIdf, "'")+1 : len(uniqueIdf)-2]
		referrer = referrer[strings.Index(referrer, "'")+1 : len(referrer)-2]
		log.Println("begTime: ", begTime)
		log.Println("uniqueIdf: ", uniqueIdf)
		log.Println("referrer: ", referrer)
		//开始前
		form := url.Values{}
		form.Add("csrf-app-name", config.User.Csrf)
		form.Add("userId", userDetailedInfo.User.UserId)
		form.Add("bookIds", bookid)
		log.Printf("POST %s\n", "https://stat.readoor.cn/index.php/materialData/getMaterialRecord")
		resp, err = client.POSTForm("https://stat.readoor.cn/index.php/materialData/getMaterialRecord", form)
		if err != nil {
			log.Println("Error: ", err)
			continue
		}
		resp.Body.Close()

		// 开始训练
		log.Printf("开始第 %d 训练 本次训练 %d s\n", i+1, t)
		time.Sleep(time.Duration(t+5) * time.Second)
		// 提交前
		form = url.Values{}
		form.Add("csrf-app-name", config.User.Csrf)
		form.Add("previousMaterialList", previousMaterialList)
		form.Add("materialRecordList", materialRecordList)
		log.Printf("POST %s\n", "https://stat.readoor.cn/index.php/materialData/insertMaterialRecord")
		resp, err = client.POSTForm("https://stat.readoor.cn/index.php/materialData/insertMaterialRecord", form)
		if err != nil {
			log.Println("Error: ", err)
			continue
		}
		resp.Body.Close()
		// 提交
		submitData = strings.Replace(submitData, "${time}", strconv.Itoa(t+rand.Int()%100-20), -1)
		submitData = strings.Replace(submitData, "${unique_idf}", uniqueIdf, -1)
		submitData = strings.Replace(submitData, "${submit_unique_num}", strconv.Itoa(int(time.Now().UnixMilli()+rand.Int64()%50)), -1)
		log.Println("submitData: ", submitData)
		form = url.Values{}
		form.Add("csrf-app-name", config.User.Csrf)
		form.Add("data", submitData)
		form.Add("cid", cid)
		form.Add("begTime", begTime)
		form.Add("uidffqbsub", uniqueIdf)
		form.Add("url", getInfoUrl)
		form.Add("referrer", referrer)
		form.Add("system", system)
		log.Printf("POST %s\n", "https://mrr.readoor.cn/api/qb/producer")
		resp, err = client.POSTForm("https://mrr.readoor.cn/api/qb/producer", form)
		if err != nil {
			log.Println("Error: ", err)
			continue
		}
		res, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error: ", err)
			continue
		}
		log.Println("res: ", string(res))

		resp.Body.Close()
		// 休息一定时间来模拟真人
		time.Sleep(time.Duration(rand.Int()%10 + 20))
	}
}
