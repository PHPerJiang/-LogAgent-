package main

import (
	"net/http"
	"sync"
)

var (
	urlChan chan string
	wrLock  sync.Mutex
	wg      sync.WaitGroup
)

func getImgUrls(url string) {
	body, err := http.Get(url)
}

func main() {

}
