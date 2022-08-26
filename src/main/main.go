package main

import (
	"go_web/src/main/form"
	"net/http"
)

func main() {
	http.HandleFunc("/", form.Escape1)
	http.ListenAndServe(":80", nil)
}
