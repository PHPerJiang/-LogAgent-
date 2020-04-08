package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/demo6", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./base.html", "./home.html", "./index.html")
		if err != nil {
			log.Printf("parsefiles failed err %v", err)
			return
		}
		t.ExecuteTemplate(w, "base.html", map[string]interface{}{
			"name": "Gopher",
		})
	})
	http.ListenAndServe(":8000", nil)
}
