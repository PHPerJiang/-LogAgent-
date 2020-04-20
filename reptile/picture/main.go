package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func GetHtmlUrls(url string) (urls []string) {
	resp, _ := http.Get(url)
	urls = make([]string, 0)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	reg := `<img[\s\S]+?src="(http[\s\S]+?)"`
	re := regexp.MustCompile(reg)
	urlsArr := re.FindAllStringSubmatch(string(body), -1)
	fmt.Println("捕获图片：", len(urlsArr), "张")
	for _, ret := range urlsArr {
		urls = append(urls, ret[1])
	}
	return urls
}

func DownImages(urls []string) {
	for _, url := range urls {
		fmt.Println(url)
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		content, _ := ioutil.ReadAll(resp.Body)
		filename := "/Users/smzdm/go/src/LogAgent/reptile/picture/images/" + strconv.Itoa(int(time.Now().UnixNano())) + ".jpg"
		err := ioutil.WriteFile(filename, content, 0644)
		if err != nil {
			fmt.Printf("下载失败, err %v", err)
		} else {
			fmt.Println("下载成功")
		}
	}
}

func main() {
	urls := GetHtmlUrls("https://www.163.com")
	DownImages(urls)
}
