package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/demo7", func(w http.ResponseWriter, r *http.Request) {
		//自定义函数 展示html文本
		showHtml := func(s string) template.HTML {
			return template.HTML(s)
		}
		//自定义符号
		t, err := template.New("demo.html").Delims("{[", "]}").Funcs(template.FuncMap{"showHtml": showHtml}).ParseFiles("./demo.html")
		if err != nil {
			log.Printf("parseFiles failed %v", err)
			return
		}
		t.Execute(w, map[string]interface{}{
			"word":    "hello Gopher!",
			"script1": "<script>alert(1);</script>",
			"script2": `<a href="https://www.baidu.com">百度</a>`,
		})
	})
	http.ListenAndServe(":8000", nil)
}
