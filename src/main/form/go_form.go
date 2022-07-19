package form

import (
	"fmt"
	"net/http"
	"strings"
)

// form参数解析
func ParamHandler(w http.ResponseWriter, r *http.Request) {
	// 解析url传递的参数，如果不执行ParseForm(),则无法取到表单数据
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key: " + k)
		fmt.Println("val: " + strings.Join(v, ""))
	}
	fmt.Fprintln(w, "parse form param success")
}

func RequestMethodHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		userName := r.Form.Get("userName")
		fmt.Fprintln(w, "Welcome "+userName)
		return
	}
	// contentType为 x-www-form-urlencoded 时可以用下面的方式取出参数
	if r.Method == "POST" {
		fmt.Println(r.PostForm)
		userName := r.PostForm["userName"]
		fmt.Println(r.PostFormValue("userName"))
		fmt.Println("userName: ", userName)
		fmt.Fprintln(w, "welcome ", userName)
		return
	}
}
