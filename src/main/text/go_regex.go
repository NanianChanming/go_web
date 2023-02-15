package text

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func RegexDemo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	strings := r.Form.Get("ip")
	fmt.Fprintf(w, "%t", IsIP(strings))
	fmt.Println("---------------------------------------------")
	GetContent()
}

/**
正则表达式是一种进行模式匹配和文本操纵的复杂而强大的工具
go语言通过regexp标准包为正则表达式提供了官方支持

通过正则判断是否匹配
regexp包中含有三个函数用来判断是否匹配，如果匹配返回true，否则返回false
func Match(pattern string, b []byte) (matched bool, error error)
func MatchReader(pattern string, r io.RuneReader) (matched bool, error error)
func MatchString(pattern string, s string) (matched bool, error error)
这三个函数都实现了同一个功能，就是判断pattern是否和输入源匹配，匹配则返回true，
如果解析出错则返回error。三个函数的输入源分别是byte slice, RuneReader和string
*/

// 判断是否为IP地址
func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

// 通过正则获取内容
func GetContent() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("http get error:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error: ", err)
		return
	}
	src := string(body)

	// 将HTML标签全部替换为小写
	reg, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	reg.ReplaceAllStringFunc(src, strings.ToLower)
	fmt.Println(strings.TrimSpace(src))

	// 去除style

}

/**
从上面的示例可以看出，使用复杂的正则首先是compile,它会解析正则表达式是否合法，如果正确，那么就会返回一个Regexp，
然后利用返回的Regexp在任意的字符串上执行需要的操作
*/
