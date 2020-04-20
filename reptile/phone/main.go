package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	reg := `1[3456789]\d{5}`
	resp, err := http.Get(`http://www.gold116.com/`)
	if err != nil {
		fmt.Println("get page failed")
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read body failed %v", err)
		return
	}
	re := regexp.MustCompile(reg)
	phoneNums := re.FindAllStringSubmatch(string(body), -1)
	fmt.Println(phoneNums)
}
