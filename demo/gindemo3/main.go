package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/demo3", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFiles("./demo3.html")
		if err != nil {
			log.Printf("parsefiles failed %v", err)
			return
		}
		tpl.Execute(w, map[string]interface{}{
			"data": map[string]interface{}{
				"data1": map[string]interface{}{
					"name":   "Gopher",
					"age":    23,
					"gender": "男",
				},
			},
		})
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Printf("listenandserver failed %v", err)
	}
}
