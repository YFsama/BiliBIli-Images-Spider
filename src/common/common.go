package common

import (
	"io/ioutil"
	"net/http"
	"os"
)



type Request struct {
	Url string
}

func Get(r Request) (body []byte, err error) {
	var (
		resp *http.Response
	)
	if resp, err = http.Get(r.Url); err != nil {
		return
	}
	//defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	resp.Body.Close()
	return body, err
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(savePath []string) (err error) {
	var (
		dirStatus bool
		dir       string
	)
	//savePath := []string{"images", "logs"}
	for _, dir = range savePath {
		if dirStatus, err = PathExists(dir); err != nil {
			continue
		}
		if !dirStatus {
			// 创建文件夹
			if err = os.Mkdir(dir, os.ModePerm); err != nil {
				return
			}
		}
	}
	return
}
