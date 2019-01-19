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
		status   bool
		pictures ApiDetailPictures
		username string
		title    string
		src      string
	)
	downOk = make(chan int)
	fmt.Println(data.User.Name + "-" + data.Item.Title + "-" + strconv.Itoa(len(data.Item.Pictures)))
	go func() {
		i = 0
		for _, pictures = range data.Item.Pictures {
			username = data.User.Name
			title = data.Item.Title
			src = pictures.ImgSrc

			go func(userName string, title string, src string,i int) {
				if err = common.CreateDir([]string{"images/" + userName}); err != nil {
					goto END
				}

				filename = "images/" + userName + "/" + title + "-" + strconv.Itoa(i) + ".jpg"
				i++
				fmt.Println(filename)
				if status, err = common.PathExists(filename); err != nil {
					goto END
				}
				if !status {
					if err = download(filename, src); err != nil {
						goto END
					}
					downOk <- 0
				}
			END:
				downOk <- 0
			}(username, title, src,i)
			i++
		}
	}()

	for range data.Item.Pictures {
		//fmt.Println(data.Item.Title)
		<-downOk
	}

}
