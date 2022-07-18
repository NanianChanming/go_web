package main

import (
	"go_web/src/main/web"
	"net/http"
)

func main() {
	http.HandleFunc("/", web.MyHandler)
	http.ListenAndServe(":8080", nil)
}
