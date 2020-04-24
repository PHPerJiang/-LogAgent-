package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var (
	urlChan chan string
	wrLock  sync.Mutex
	wg      sync.WaitGroup
	regs    = map[string]string{
		"imgUrl": `<img[\s\S]+?src=['"](http[^'"]+?\.(?:jpg|jpeg|png))[^>]*>`,
	}
)

//处理错误信息
func handleEerr(err error) {
	if err != nil {
		log.Printf("error happend! err:%v\n", err)
		return
	}
}

//获取图片写入队列
func getImgUrls(url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	handleEerr(err)
	body, err := ioutil.ReadAll(resp.Body)
	handleEerr(err)
	re := regexp.MustCompile(regs["imgUrl"])
	urlSlice := re.FindAllStringSubmatch(string(body), -1)
	for _, v := range urlSlice {
		urlChan <- v[1]
	}
}

//下载图片
func downImg(imgUrl string) {
	resp, err := http.Get(imgUrl)
	defer resp.Body.Close()
	handleEerr(err)
	body, err := ioutil.ReadAll(resp.Body)
	handleEerr(err)
	filename := "/Users/smzdm/go/src/LogAgent/reptile/async_picture/images/" + strconv.Itoa(int(time.Now().Nanosecond())) + ".jpg"
	err = ioutil.WriteFile(filename, body, 0644)
	handleEerr(err)
	log.Println(filename + " download success!")
}

func main() {
	urlChan = make(chan string, 10)
	go getImgUrls("https://www.nvshens.net/g/32446/8.html")
	for {
		select {
		case url := <-urlChan:
			log.Println(url)
			downImg(url)
		default:
			log.Println("wait...")
			time.Sleep(time.Second)
		}
	}
}
