package bilibili

import (
	"fmt"
	"github.com/yranarf/BiliBIli-Images-Spider/src/common"
	"io"
	"net/http"
	"os"
	"strconv"
)

func InitSpider(url string) (err error) {
	err = indexSpider(url)
	return
}

//TODO 下载优化
func download(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func detailDownload(data ApiDetailData) {
	var (
		err      error
		filename string
		i        int
		pictures ApiDetailPictures
	)
	if err = common.CreateDir([]string{"images/" + data.User.Name}); err != nil {
		return
	}
	for i = 0; i >= len(data.Item.Pictures); i ++ {

	}
	i = 0
	for _, pictures = range data.Item.Pictures {
		filename = "images/"+ data.User.Name + "/" + data.Item.Title + "-" + strconv.Itoa(i) + ".jpg"
		if err = download(filename, pictures.ImgSrc); err != nil {
			return
		}
		i++
	}

}
