package go_gin

import (
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
