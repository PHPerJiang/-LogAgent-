package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func GetHtml(url string) (string, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
func main() {
	body, err := GetHtml("https://www.163.com")
	if err != nil {
		fmt.Printf("get bofy failed %v", err)
		return
	}
	reg := `<img[\s\S]+?src="(http[\s\S]+?)"`
	re := regexp.MustCompile(reg)
	urls := re.FindAllStringSubmatch(body, -1)
	for _, ret := range urls {
		fmt.Println(ret[1])
	}
}
