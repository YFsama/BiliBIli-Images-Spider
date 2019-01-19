package bilibili

import (
	"encoding/json"
	"github.com/yranarf/BiliBIli-Images-Spider/src/common"
	"strconv"
	"time"
)

type ApiListResponse struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    ApiListData `json:"data"`
	Message string      `json:"message"`
}

type ApiListData struct {
	TotalCount int            `json:"total_count"`
	Items      []ApiListItems `json:"items"`
}
type ApiListItems struct {
	User ApiListUser `json:"user"`
	Item ApiListItem `json:"item"`
}

type ApiListItem struct {
	DocId        int               `json:"doc_id"`
	PosterUid    int               `json:"poster_uid"`
	Pictures     []ApiListPictures `json:"pictures"`
	Title        string            `json:"title"`
	Category     string            `json:"category"`
	UploadTime   int               `json:"upload_time"`
	AlreadyLiked int               `json:"already_liked"`
	AlreadyVoted int               `json:"already_voted"`
}

type ApiListPictures struct {
	ImgSrc    string `json:"img_src"`
	ImgWidth  int    `json:"img_width"`
	ImgHeight int    `json:"img_height"`
	ImgSize   int    `json:"img_size"`
}

type ApiListUser struct {
	Uid     int    `json:"uid"`
	HeadUrl string `json:"head_url"`
	Name    string `json:"name"`
}

var listOk chan int
var downOk chan int
var docListOk chan int
var detailOk chan int

func indexSpider(area string) (err error) {
	var (
		body       []byte
		apiResp    ApiListResponse
		i          int
		request    common.Request
		requestUrl string
		url        string
		quitI      int
	)

	listOk = make(chan int)
	url = "https://api.vc.bilibili.com/link_draw/v2/" + area + "/list"
	//0~24
	for i = 0; i <= 24; i++ {
		//time.Sleep(50 * time.Millisecond)
		go func(i int) {
			requestUrl = url + "?category=cos&type=hot&page_num=" + strconv.Itoa(i) + "&page_size=20"
			//fmt.Println(requestUrl)
			request = common.Request{
				Url: requestUrl,
			}
			if body, err = common.Get(request); err != nil {
				goto END
			}

			if err = json.Unmarshal(body, &apiResp); err != nil {
				goto END
			}

			if apiResp.Code == 0 && apiResp.Msg == "success" {
				apiListResponse(apiResp)
				listOk <- 0

			}
		END:
			listOk <- 0
		}(i)
	}

	for quitI = 0; quitI <= 1; quitI++ {
		<-listOk
	}

	err = nil
	return
}



func apiListResponse(resp ApiListResponse) {
	var (
		docList   []int
		request   common.Request
		detailUrl string
		items     ApiListItems
		doc       int
		quitI     int
	)
	docList = make([]int, 0)
	for _, items = range resp.Data.Items {
		docList = append(docList, items.Item.DocId)
	}

	docListOk = make(chan int)
	detailOk = make(chan int)
	for _,doc = range docList{
		go func(doc int) {
			detailUrl = "https://api.vc.bilibili.com/link_draw/v1/doc/detail?doc_id=" + strconv.Itoa(doc)
			//fmt.Println(doc)
			request = common.Request{
				Url: detailUrl,
			}
			go detailSpider(request)
			docListOk <- 0
		}(doc)
	}

	for quitI = 0; quitI <= len(docList); quitI++ {
		<-docListOk
	}
	for range docList{
		<-detailOk
	}
time.Sleep(3 * time.Second)
	return
}
