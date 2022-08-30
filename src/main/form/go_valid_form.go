package form

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

/**
表单验证
*/

func ValidForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// 必填
	if len(r.Form.Get("username")) == 0 {
		fmt.Fprintf(w, "username can't be null")
	}
	// 数字校验
	// 整数，例如年龄
	atoi, err := strconv.Atoi(r.Form.Get("age"))
	if err != nil {
		fmt.Fprintf(w, "年龄格式输入错误")
	}
	if atoi > 100 {
		fmt.Fprintf(w, "年龄值不合理")
	}
	// 正则表达式
	// 中文 英文
	if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("cnname")); !m {
		fmt.Fprintf(w, "中文名称不对")
	}
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("enname")); !m {
		fmt.Fprintf(w, "英文名不对")
	}
}
