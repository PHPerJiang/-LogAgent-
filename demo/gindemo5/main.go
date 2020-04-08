package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/demo5", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./tpl1.html", "./ul.tmpl")
		if err != nil {
			log.Printf("parsefiles failed err %v", err)
			return
		}
		t.Execute(w, map[string]interface{}{
			"name": "hello Gopher",
		})
	})
	http.ListenAndServe(":8000", nil)
}
