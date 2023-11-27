package go_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"html/template"
	"log"
	"net/http"
)

/*
InitGin
gin 快速入门
1.安装gin:
go get -u github.com/gin-gonic/gin
*/
func InitGin() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "success"})
	})
	r.GET("/asciiJson", AsciiJSON)
	// 监听并且在0.0.0.0:8080上启动服务
	r.Run(":8080")
}

/*
AsciiJSON
使用AsciiJSON生成具有专一ASCII字符的ASCII-only JSON
*/
func AsciiJSON(ctx *gin.Context) {
	data := map[string]interface{}{
		"lang": "Go",
		"tag":  "<br>",
	}
	ctx.AsciiJSON(http.StatusOK, data)
}

/*
HtmlDemo
使用LoadHTMLGlob()或者LoadHTMLFiles()
*/
func HtmlDemo() {
	router := gin.Default()
	router.LoadHTMLGlob("./src/templates/*")
	//router.LoadHTMLFiles("templates/index.html", "templates/index1.html")
	router.GET("index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello World",
		})
	})
	router.Run()
}

var html = template.Must(template.New("http2Push").Parse(`
	<html>
	<head>
	  <title>Https Test</title>
	  <script src="/templates/app.js"></script>
	</head>
	<body>
	  <h1>Welcome, Ginner!</h1>
	</body>
	</html>
	`))

func Http2Push() {
	router := gin.Default()
	//router.Static("/templates", "./templates")
	//router.SetHTMLTemplate(html)

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "This is the main page!")
		if push, ok := context.Writer.(http.Pusher); ok {
			if err := push.Push("./src/templates/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	http2.ConfigureServer(server, nil)
	server.ListenAndServe()
}

/*
JSONP
使用jsonp向不同域的服务器请求数据，如果查询参数存在回调，则将回调添加到响应体中
*/
func JSONP() {
	router := gin.Default()
	router.GET("/jsonp", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		context.JSONP(http.StatusOK, data)
	})
	router.Run()
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

/*
MultipartUrlencodedFormParams
Multipart/Urlencoded 绑定
(请求参数绑定)
*/
func MultipartUrlencodedFormParams() {
	router := gin.Default()
	router.POST("/login", func(context *gin.Context) {
		var form LoginForm
		// 可以使用显式绑定声明绑定 Multipart form
		// context.ShouldBindWith(&form, binding.Form)
		// 或者简单的使用ShouldBind 方法自动绑定
		// 在这种情况下将选择合适的绑定
		if context.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				context.JSON(http.StatusOK, gin.H{"msg": "login success"})
			} else {
				context.JSON(http.StatusUnauthorized, gin.H{"msg": "login failed"})
			}
		}
	})
	router.Run()
}

/*
MultipartUrlencodedForm
表单绑定
*/
func MultipartUrlencodedForm() {
	router := gin.Default()
	router.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous")
		context.JSON(http.StatusOK, gin.H{
			"msg":     "success",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run()
}

/*
PureJson
通常，json使用unicode替换HTML字符，例如<变为\u003c,如果按字面对这些字符进行编码，则可以使用PureJson，go1.6及更低版本无法使用此功能
*/
func PureJson() {
	router := gin.Default()
	router.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello, World!</b>",
		})
	})
	router.GET("/purejson", func(context *gin.Context) {
		context.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello, World!</b>",
		})
	})
	router.Run()
}

/*
QueryPostForm
POST /post?id=1234&page=1
HTTP/1.1 Content-Type: application/x-www-form-urlencoded
*/
func QueryPostForm() {
	router := gin.Default()
	router.POST("/query_post_form", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.PostForm("message")

		fmt.Printf("id: %s, page: %s, name: %s, message: %s \n", id, page, name, message)
	})
	router.Run()
}
