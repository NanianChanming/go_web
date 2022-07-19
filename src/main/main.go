package main

import (
	"go_web/src/main/form"
	"net/http"
)

func main() {
	http.HandleFunc("/", form.ParamHandler)
	http.HandleFunc("/login", form.RequestMethodHandler)
	http.ListenAndServe(":80", nil)
}
