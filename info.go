package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/woshikedayaa/ixue_note/internal/models"
	"github.com/woshikedayaa/ixue_note/internal/utils"
	"io"
	"log"
	"net/http"
	"net/url"
)

var TargetBookInfo models.BundleBook

func GetBooKList() error {
	availableBookList, err := getBookAvailableList()
	if err != nil {
		return err
	}

	if len(availableBookList) == 0 {
		return errors.New("no book available")
	}
	detailedInfo, err := getBookDetailedInfo(availableBookList)
	if err != nil {
		return err
	}
	if len(detailedInfo) != len(availableBookList) {
		log.Println("注意: 获取到详细关于书的信息与可用书列表的数量不同 可能有请求失败")
		log.Printf("可用的书 %d , 实际获取了 %d \n", len(availableBookList), len(detailedInfo))
	}
	targetBookIdx := 0 // default 0
	for i := 0; i < len(detailedInfo); i++ {
		fmt.Printf("有以下书 %d : %s", i, detailedInfo[i].Data.BookName)
	}
	fmt.Printf("请输入一个数字来选择书:")
	for targetBookIdx < len(detailedInfo) {
		_, err = fmt.Scanf("%d", &targetBookIdx)
		if err != nil {
			return err
		}
		if targetBookIdx >= len(detailedInfo) {
			log.Println("无效的数字")
			targetBookIdx = 0
			continue
		} else {
			break
		}
	}
	// bind info
	TargetBookInfo = models.BundleBook{}

	TargetBookInfo.SimpleBook = availableBookList[targetBookIdx]
	TargetBookInfo.DetailedBook = detailedInfo[targetBookIdx]
	return nil
}

func getBookAvailableList() (availableBookList []models.SimpleBook, err error) {
	var (
		bookListUrl = utils.URLBuild("https://app.readoor.cn/app/dt/ln/1544059443/11").AddArg("s", "11")
		resp        *http.Response
		HTMLdoc     *goquery.Document
	)

	log.Printf("GET %s\n", bookListUrl.Build())
	resp, err = client.GET(bookListUrl.Build())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	HTMLdoc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	HTMLdoc.Find(".data_bind_dom").Each(func(i int, selection *goquery.Selection) {
		book := models.SimpleBook{}
		book.BookIDEncrypt, _ = selection.Attr("data-book_id")
		book.Href, _ = selection.Attr("data-href")
		book.Pid, _ = selection.Attr("data-pid")
		book.GroupID, _ = selection.Attr("data-group_id")
		book.TrainID, _ = selection.Attr("data-training_id")
		book.Type, _ = selection.Attr("data-type")
		availableBookList = append(availableBookList, book)
	})
	return availableBookList, nil
}

func getBookDetailedInfo(availableBookList []models.SimpleBook) ([]models.DetailedBook, error) {
	// log
	log.Println("找到可用的书数量:", len(availableBookList))
	for i := 0; i < len(availableBookList); i++ {
		log.Println(i, ":", availableBookList[i].BookIDEncrypt)
	}
	// request
	var (
		err  error
		resp *http.Response

		detailedBookList []models.DetailedBook
		detailedBookUrl  = "https://app.readoor.cn/app/dm/getTtclassInfo/1544059443"
		detailedBookResp []byte
	)

	for i := 0; i < len(availableBookList); i++ {
		book := models.DetailedBook{}
		log.Printf("POST %s for %s\n", detailedBookUrl, availableBookList[i].BookIDEncrypt)
		data := url.Values{}
		data.Add("csrf_app_name", config.User.Csrf)
		data.Add("book_id", availableBookList[i].BookIDEncrypt)
		resp, err = client.POSTForm(detailedBookUrl, data)
		if err != nil {
			return nil, err
		}
		detailedBookResp, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(detailedBookResp, &book)
		if err != nil {
			return nil, err
		}
		detailedBookList = append(detailedBookList, book)
		// close
		resp.Body.Close()
	}
	return detailedBookList, nil
}
