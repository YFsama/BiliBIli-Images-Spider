package storage

import (
	"github.com/yranarf/BiliBIli-Images-Spider/src/common"
)

const IMAGES = "images"
const LOGS  = "logs"

func initLocal() (err error) {
	err = common.CreateDir([]string{"images", "logs"})

	return
}


