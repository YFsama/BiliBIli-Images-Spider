package worker

import "github.com/yranarf/BiliBIli-Images-Spider/src/bilibili"

func InitWorker(url string){
	var (
		err error
	)
	if err = bilibili.InitSpider(url);err != nil{
		panic(err)
	}
}


