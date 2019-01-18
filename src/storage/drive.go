package storage

import (
	"fmt"
	"github.com/yranarf/BiliBIli-Images-Spider/src/bilibili"
)


func InitStorage(drive string) {
	var (
		err error
	)
	switch drive {
	case "local":
		err = initLocal()
		break
	}

	//for {
	//
	//}

	if err != nil {
		fmt.Println(err)
	}
}

func SaveItem(data bilibili.ApiDetailData) {
	fmt.Println(data.User.Name)
	//var (
	//	err error
	//)
	//fmt.Println(item.Item.Title)
	//if err = common.CreateDir([]string{IMAGES + "\\" + item.User.Name}); err != nil {
	//
	//}
	//
	//return
}
