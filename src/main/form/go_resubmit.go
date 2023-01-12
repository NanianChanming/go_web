package form

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// 防止表单重复提交
/*
教程中给出的方案为:
增加一个隐藏字段当作唯一标识，提交到后端后进行存储，
二次提交的时候进行查找此值来决定是否为重复提交
*/

func ValidResubmit(w http.ResponseWriter, r *http.Request) {
	// 以登录为例
	if r.Method == "GET" {
		// 请求登录的页面 写回给前端token（token生成方式: md5当前时间戳）
		unix := time.Now().Unix()
		hash := md5.New()
		io.WriteString(hash, strconv.FormatInt(unix, 10))
		token := fmt.Sprintf("%x", hash.Sum(nil))
		fmt.Fprintln(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// 验证 token 的合法性
		} else {
			// token为空则报错
		}
	}
}
