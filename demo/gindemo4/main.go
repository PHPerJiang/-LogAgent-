package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/demo4", func(w http.ResponseWriter, r *http.Request) {
		f := func(name string) string {
			return "hello ~ " + name
		}
		t, err := template.New("demo4.html").Funcs(template.FuncMap{
			"addName": f,
		}).ParseFiles("./demo4.html")
		if err != nil {
			log.Printf("parseFiles failed %v", err)
			return
		}
		t.Execute(w, map[string]interface{}{
			"name": "Gopher",
		})
	})
	http.ListenAndServe(":8000", nil)
}
