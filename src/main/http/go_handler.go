package http

import (
	"fmt"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.Form.Get("userName")
	fmt.Fprintln(w, "hello my name is "+userName)
}
