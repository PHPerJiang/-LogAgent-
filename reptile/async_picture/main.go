package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var (
	urlChan  chan string
	pageChan chan string
	wg       sync.WaitGroup
	regs     = map[string]string{
		"imgUrl":  `<img[\s\S]+?src=['"](http[^'"]+?\.(?:jpg|jpeg|png))[^>]*>`,
		"pageUrl": `<a[\s\S]+?href="(https://wiki.smzdm.com/kongtiao/p[\s\S]+?\/)"`,
	}
	pageMap  map[string]string
	imageMap map[string]string
	taskNum  int = 10
)

func init() {
	urlChan = make(chan string, 10000)
	pageChan = make(chan string, 100)
	pageMap = make(map[string]string, 0)
	imageMap = make(map[string]string, 0)
}

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
	//匹配页面中的图片
	re := regexp.MustCompile(regs["imgUrl"])
	urlSlice := re.FindAllStringSubmatch(string(body), -1)
	for _, v := range urlSlice {
		urlChan <- v[1]
	}
	fmt.Println("共抓取到" + strconv.Itoa(len(urlChan)) + "张图片")
	//匹配页面中的pageurl
	re = regexp.MustCompile(regs["pageUrl"])
	pageUrlSlice := re.FindAllStringSubmatch(string(body), -1)
	for _, v := range pageUrlSlice {
		if _, ok := pageMap[v[1]]; !ok {
			pageMap[v[1]] = v[1]
			pageChan <- v[1]
		}
	}
	fmt.Println("共抓取到" + strconv.Itoa(len(pageChan)) + "个其他分页")
	wg.Done()
}

//
func getPage() {
	for {
		select {
		case pageUrl := <-pageChan:
			wg.Add(1)
			go getImgUrls(pageUrl)
		default:
			time.Sleep(time.Second)
		}
	}
}

//下载图片
func downImg(taskID int) {
	fmt.Printf("task %d start...\n", taskID)
	for {
		if url, ok := <-urlChan; !ok {
			break
		} else {
			resp, err := http.Get(url)
			defer resp.Body.Close()
			handleEerr(err)
			body, err := ioutil.ReadAll(resp.Body)
			handleEerr(err)
			filename := "/Users/smzdm/go/src/LogAgent/reptile/async_picture/images/taskID:" + strconv.Itoa(taskID) + "&time:" + strconv.Itoa(int(time.Now().Nanosecond())) + ".jpg"
			err = ioutil.WriteFile(filename, body, 0644)
			handleEerr(err)
			log.Println(filename + " download success!")
		}
	}
	fmt.Printf("task %d end...\n", taskID)
	wg.Done()
}

//起多个协程并发下载
func multiDown() {
	for i := 1; i <= taskNum; i++ {
		wg.Add(1)
		go downImg(i)
	}
}

func main() {
	wg.Add(1)
	getImgUrls("https://wiki.smzdm.com/kongtiao/")
	go getPage()
	multiDown()
	wg.Wait()
}
