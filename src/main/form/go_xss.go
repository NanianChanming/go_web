package form

import (
	"fmt"
	"html"
	"net/http"
)

/**
web网站包含大量动态内容，所谓动态内容，就是根据用户环境需要，web应用程序能够输出相应的内容，
动态网站容易受到 跨站脚本攻击 Cross Site Scripting 通常缩写为XSS， 而静态网站则不受影响
攻击者通常会在有漏洞的程序中插入JavaScript、VBScript、ActiveX或Flash以欺骗用户。

对XSS最佳的防护应该结合以下两种方法：
一是验证所有输入数据，有效检测攻击；另一个是对所有输出数据进行适当的处理，以防止任何已经成功注入的脚本在浏览器运行
*/
/**
Go语言的html/template里面带有下面几个函数可以帮我们转义
func HTMLEscape(w io.Writer, b[] byte) //把b进行转移之后写到w
func HTMLEScapeString(s string) string //转义s之后返回结果字符串
func HTMLEscapeer(args ...interface{}) string //支持多个参数一起转义，返回结果字符串
*/

func Escape1(w http.ResponseWriter, r *http.Request) {
	//传入参数：username
	//假如我们输入的username是 <script>alert()</script> 那么浏览器就会弹出一个空白框
	r.ParseForm()
	username := r.Form.Get("username")
	fmt.Println("username = ", username)
	//转义之后浏览器上会看到: &lt;script&gt;alert()&lt;/script&gt;
	username = html.EscapeString(username)
	fmt.Println("escape username = ", username)
	fmt.Fprintln(w, username)
}
