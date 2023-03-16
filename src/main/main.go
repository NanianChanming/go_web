package main

import (
	_ "go_web/src/main/web"
	"net/http"
)

func main() {
	//http.HandleFunc("/socket", web.WebSocketHandle)
	http.ListenAndServe(":8080", nil)
}

//var Route map[string]interface{}
//
//func init() {
//	Route["/generateJson"] = text.GenerateJson
//}
//
//func getRoute(url string) interface{} {
//	return Route[url]
//}
