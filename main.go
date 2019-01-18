package main

import (
	"flag"
	"github.com/yranarf/BiliBIli-Images-Spider/src/storage"
	"github.com/yranarf/BiliBIli-Images-Spider/src/worker"
)

var (
	area         string
	storageDrive string
)

func initArea() () {
	flag.StringVar(&area, "area", "photo", "请输入要爬取的分区")
	flag.StringVar(&storageDrive, "storage", "local", "请输入要使用存储驱动")
}

func main() {
	initArea()
	flag.Parse()

	storage.InitStorage(storageDrive)
	worker.InitWorker(area)
}
