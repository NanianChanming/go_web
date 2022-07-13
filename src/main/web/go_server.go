package web

import (
	"fmt"
	"net/http"
)

func HelloWorldServer(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello World")
}
