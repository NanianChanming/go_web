package security

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

/*
CSRF（Cross-site request forgery）,中文名称：跨站请求伪造，
也被称为：one click attack/session riding 缩写为：CSRF/XSRF。
如何预防CSRF
CSRF的防御可以从服务端和客户端两方面着手，防御效果是从服务端着手效果比较好，现在一般的CSRF防御也在服务端进行。
服务端的预防CSRF攻击的方式方法有多种，但思想上都是差不多的，主要从以下两个方面入手
·正确使用GET，POST和Cookie
·在非GET请求中增加伪随机数
以go语言举例，如何限制对资源的访问方法
mux.Get("/user/:uid", getUser)
mux.Post("/user/:uid", modifyUser)
这样处理后，因为限定了只有POST能修改，当GET方式请求时就拒绝响应，所以GET方式的CSRF攻击就可以防止了，
但这样不能全部解决问题，因为POST也可以模拟。
因此需要实施第二步，在非GET方式的请求中增加随机数，这个大概有三种方式来进行
·为每个用户生成一个唯一的cookie token，所有表单都包含同一个伪随机值，这种方案最简单，
因为攻击者不能获得第三方的Cookie（理论上），所以表单中的数据也就构造失败，
但由于用户的Cookie很容易由于网站的XSS漏洞而被盗取，所以这个方案必须要在没有XSS的情况下才安全。
·每个请求使用验证码，这个方案是完美的，但因为要多次输入验证码，所以用户友好性很差，所以不适合实际应用。
·不同的表单包含一个不同的伪随机值，实现如下：
*/
func randomToken(w http.ResponseWriter) {
	hash := md5.New()
	io.WriteString(hash, strconv.FormatInt(time.Now().Unix(), 10))
	io.WriteString(hash, "xxxxxxxxxxxxxxxx")
	token := fmt.Sprintf("%x", hash.Sum(nil))

	t, _ := template.ParseFiles("login.gtpl")
	t.Execute(w, token)
}

/*
验证token
*/
func verifyToken(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := r.Form.Get("token")
	if token != "" {
		// 验证token合法性
	} else {
		// 不存在token则报错
	}
}

/*
这样基本就实现了安全的POST。

总结：
跨站请求伪造，即CSRF， 是一种非常危险的Web安全威胁，它被Web安全界成为“沉睡的巨人”，其威胁成都由此“美誉”便可见一斑。
*/
