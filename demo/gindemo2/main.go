package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/gindemo2/demo2", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./demo1.html")
		if err != nil {
			log.Println("parse files failed")
			return
		}
		var m = map[string]interface{}{
			"name":   "Gopher",
			"age":    23,
			"gender": "男",
		}
		t.Execute(w, m)
	})
	http.ListenAndServe(":8000", nil)
}
