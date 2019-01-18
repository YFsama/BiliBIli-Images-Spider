package bilibili

import (
	"encoding/json"
	"github.com/yranarf/BiliBIli-Images-Spider/src/common"
	"strconv"
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

func indexSpider(area string) (err error) {
	var (
		body       []byte
		apiResp    ApiListResponse
		i          int
		request    common.Request
		requestUrl string
		url        string
	)
	url = "https://api.vc.bilibili.com/link_draw/v2/" + area + "/list"
	//0~24
	func() {
		for i = 0; i <= 1; i++ {
			requestUrl = url + "?category=cos&type=hot&page_num=" + strconv.Itoa(i) + "&page_size=20"

			request = common.Request{
				Url: requestUrl,
			}
			if body, err = common.Get(request); err != nil {
				continue
			}

			if err = json.Unmarshal(body, &apiResp); err != nil {
				continue
			}

			if apiResp.Code == 0 && apiResp.Msg == "success" {
				apiListResponse(apiResp)
			} else {
				continue
			}
		}

	}()

	return
}

func apiListResponse(resp ApiListResponse) {
	var (
		docList   []int
		request   common.Request
		detailUrl string
		items     ApiListItems
		doc       int
	)
	docList = make([]int, 0)
	for _, items = range resp.Data.Items {
		docList = append(docList, items.Item.DocId)
	}
	//fmt.Println(items.Item.Title)
	for _, doc = range docList {
		detailUrl = "https://api.vc.bilibili.com/link_draw/v1/doc/detail?doc_id=" + strconv.Itoa(doc)
		request = common.Request{
			Url: detailUrl,
		}
		 detailSpider(request)
	}

}
