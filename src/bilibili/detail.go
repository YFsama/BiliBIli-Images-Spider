package bilibili

import (
	"encoding/json"
	"fmt"
	"github.com/yranarf/BiliBIli-Images-Spider/src/common"
)

type ApiDetailResponse struct {
	Code    int           `json:"code"`
	Msg     string        `json:"msg"`
	Message string        `json:"message"`
	Data    ApiDetailData `json:"data"`
}

type ApiDetailData struct {
	User ApiDetailUser `json:"user"`
	Item ApiDetailItem `json:"item"`
}

type ApiDetailUser struct {
	Uid         int    `json:"uid"`
	HeadUrl     string `json:"head_url"`
	Name        string `json:"name"`
	UploadCount int    `json:"upload_count"`
}

type ApiDetailItem struct {
	Biz             int                 `json:"biz"`
	DocId           int                 `json:"doc_id"`
	PosterDoc       int                 `json:"poster_doc"`
	Category        string              `json:"category"`
	Type            int                 `json:"type"`
	Title           string              `json:"title"`
	Pictures        []ApiDetailPictures `json:"pictures"`
	UploadTime      string              `json:"upload_time"`
	UploadTimestamp int                 `json:"upload_timestamp"`
}

type ApiDetailPictures struct {
	ImgSrc    string `json:"img_src"`
	ImgWidth  int    `json:"img_width"`
	ImgHeight int    `json:"img_height"`
	ImgSize   int    `json:"img_size"`
}

func detailSpider(r common.Request) {
	var (
		body          []byte
		err           error
		apiDetailResp ApiDetailResponse
	)

	if body, err = common.Get(r); err != nil {
		return
	}

	if err = json.Unmarshal(body, &apiDetailResp); err != nil {
		return
	}

	fmt.Println(apiDetailResp.Data.Item.Title)

	if apiDetailResp.Code == 0 && apiDetailResp.Msg == "success" {
		detailDownload(apiDetailResp.Data)
	} else {
		return
	}
}
