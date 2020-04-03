package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func demo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1>", "hello word")
}

func demo1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./demo.htm")
	if err != nil {
		log.Println("read file failed")
		return
	}
	fmt.Fprintln(w, string(b))
}

func main() {
	// http.HandleFunc("/httpdemo", demo)
	http.HandleFunc("/httpdemo", demo1)
	http.ListenAndServe(":8000", nil)
}
