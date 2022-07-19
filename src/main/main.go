package main

import (
	"go_web/src/main/form"
	"net/http"
)

func main() {
	http.HandleFunc("/", form.ParamHandler)
	http.HandleFunc("/login", form.RequestMethodHandler)
	http.HandleFunc("/body", form.RequestParamBody)
	http.ListenAndServe(":80", nil)
}
