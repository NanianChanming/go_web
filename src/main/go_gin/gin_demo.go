package go_gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"golang.org/x/net/http2"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

/*
SecureJson
使用securejson防止json劫持，如果给定的结构是数组值，则默认预置"while(1),"到响应体
*/
func SecureJson() {
	router := gin.Default()
	router.GET("/secureJson", func(context *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		// 将输出：while(1);["lena", "austin", "foo"]
		context.SecureJSON(http.StatusOK, names)
	})
	router.Run()
}

/*
ConfigFile
xml/json/yaml/protobuf渲染
*/
func ConfigFile() {
	r := gin.Default()
	// gin.H是map[string]interface{}的一种快捷方式
	r.GET("/someJson", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	r.GET("/moreJson", func(context *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意msg.Name在json中定义成了user
		// 将输出：{"user": "Lena"}
		context.JSON(http.StatusOK, msg)
	})

	r.GET("someXML", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"message": "hey xml", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"message": "hey yaml", "status": http.StatusOK})
	})

	r.GET("/someProtoBuf", func(context *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protoBuf 的具体定义写在testdata/protoexample文件中
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被protoexample.Test protobuf 序列化了的数据
		context.ProtoBuf(http.StatusOK, data)
	})

	r.Run()
}

/*
UploadFile
单文件上传
*/
func UploadFile() {
	router := gin.Default()
	// 为multipart forms 设置较低的内存限制(默认是32MiB)
	router.MaxMultipartMemory = 8 << 20 // 8MiB
	router.POST("/uploadFile", func(context *gin.Context) {
		// 单文件
		file, _ := context.FormFile("file")
		log.Println(file.Filename)
		dst := "./file/" + file.Filename
		// 上传文件至指定的完整文件路径
		context.SaveUploadedFile(file, dst)
		context.JSON(http.StatusOK, gin.H{"status": "success!"})
	})
	router.Run()
}

/*
BatchUpload
多文件批量上传
*/
func BatchUpload() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/batchUpload", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
			dst := "./file/" + file.Filename
			context.SaveUploadedFile(file, dst)
		}
		context.JSON(http.StatusOK, gin.H{"status": "success!"})
	})
	router.Run()
}

/*
Reader
从reader读取数据
*/
func Reader() {
	router := gin.Default()
	router.GET("/someDataFromReader", func(context *gin.Context) {
		resp, err := http.Get("https://img0.baidu.com/it/u=3028168707,3962278789&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=281")
		if err != nil || resp.StatusCode != http.StatusOK {
			context.Status(http.StatusServiceUnavailable)
			return
		}
		reader := resp.Body
		contentLength := resp.ContentLength
		contentType := resp.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"content-Dispostion": `attachment;filename=gopher.jpg`,
		}
		context.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	router.Run()
}

/*
Restarted
优雅的重启或停止
*/
func Restarted() {
	router := gin.Default()
	router.GET("/shutdown", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置5秒地超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123456"},
	"austin": gin.H{"email": "austin@austin.com", "phone": "22222"},
	"lena":   gin.H{"email": "lena@lena.com", "phone": "33333"},
}

/*
BasicAuth
BasicAuth中间件
*/
func BasicAuth() {
	r := gin.Default()
	// 路由组使用gin.BasicAuth中间件
	// gin.Accounts是map[string]string的一种快捷方式
	authrorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	// /admin/secrets端点
	authrorized.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "no secret"})
		}
	})
	r.Run(":8080")
}

/*
HttpFunc
使用http方法
*/
func HttpFunc() {
	// 使用默认中间件（logger和recovery中间件）创建gin路由
	router := gin.Default()

	router.GET("")
	router.POST("")
	router.PUT("")
	router.DELETE("")
	router.PATCH("")
	router.HEAD("")
	router.OPTIONS("")

	router.Run()
}

/*
Middleware
使用中间件
*/
func Middleware() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入gin.DefaultWriter, 即使将GIN_HOME设置为release，
	// By Default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件会recover任何panic，如果有panic的话，会写入500
	r.Use(gin.Recovery())

	// 你可以为每个路由添加任意数量的中间件
	r.GET("/benchmark", gin.Logger())

	// 认证路由组
	// authorized := r.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样
	// authorized := r.Group("/")
	// 路由组中间件 在此例中，我们在“authorized”路由组中使用自定义创建的AuthRequired()中间件
	/*authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)
		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}*/

	r.Run()
}

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

/*
BindingURLParam
只绑定url查询字符串
ShouldBindQuery函数只绑定url查询参数而忽略post数据
*/
func BindingURLParam() {
	router := gin.Default()
	router.Any("/testing", bindUrlParam)
	router.Run()
}

func bindUrlParam(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("Only Bind By Query string")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "success")
}

/*
GoroutineInMiddleWare
在中间件中使用goroutine
当在中间件或handler中启动新的goroutine时，不能使用原始的上下文，必须使用只读副本
*/
func GoroutineInMiddleWare() {
	router := gin.Default()
	router.GET("/long_async", func(c *gin.Context) {
		// 创建在goroutine中使用的副本
		c2 := c.Copy()
		go func() {
			// 用time.Sleep() 模拟一个长任务
			time.Sleep(5 * time.Second)
			// 请注意您使用的是复制的上下文c2, 这一点很重要
			log.Println("Done! in path " + c2.Request.URL.Path)
		}()
	})
	router.GET("/long_sync", func(c *gin.Context) {
		// 用time.Sleep模拟一个长任务
		time.Sleep(5 * time.Second)
		// 因为没有使用goroutine，不需要copy上下文
		log.Println("Done! in path " + c.Request.URL.Path)
	})
	router.Run()
}

/*
LogPrint
记录日志
*/
func LogPrint() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色
	gin.DisableConsoleColor()
	// 记录到文件
	file, _ := os.Create("/data/gin.log")
	gin.DefaultWriter = io.MultiWriter(file)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码
	//gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.Run()
}

/*
LogFormat
定义路由日志的格式
*/
func LogFormat() {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	router.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	router.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	router.Run()
}

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func RequestBody() {
	router := gin.Default()
	router.POST("/requestBody", RequestBodyHandler)
	router.POST("/requestBody2", RequestBodyHandlerMuti)
	router.Run()
}

/*
RequestBodyHandler
一般通过调用c.Request.Body方法绑定数据，但不能多次调用这个方法
*/
func RequestBodyHandler(c *gin.Context) {
	a := formA{}
	b := formB{}
	// c.ShouldBind 使用了c.Request.Body 不可重用
	if errA := c.ShouldBind(&a); errA == nil {
		c.String(http.StatusOK, "the body should be formA")
		// 现在c.Request.Body 是EOF, 所以这里会报错
	} else if errB := c.ShouldBind(&b); errB == nil {
		c.String(http.StatusOK, "the body should be formB")
	}
}

/*
RequestBodyHandlerMuti
要想多次绑定，可以使用c.ShouldBindBodyWith
*/
func RequestBodyHandlerMuti(c *gin.Context) {
	a := formA{}
	b := formB{}
	// 读取c.Request.Body并将结果存入上下文
	if errA := c.ShouldBindBodyWith(&a, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
	} else if errB := c.ShouldBindBodyWith(&b, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	}
}

/*
LogColor
控制日志输出颜色
*/
func LogColor() {
	// 禁止日志颜色
	//gin.DisableConsoleColor()

	// 强制日志颜色化
	gin.ForceConsoleColor()

	// 用默认中间件创建一个gin路由
	// 日志和恢复中间件
	router := gin.Default()
	router.GET("/logColor", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})
	router.Run()
}

/*
LetsEncrypt
一行代码支持letsEncrypt https servers 实例
*/
func LetsEncrypt() {
	router := gin.Default()
	// ping handler
	router.GET("/letsEncrypt", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})
	log.Fatal(autotls.Run(router, "example1.com", "example2.com"))
}

/*
UrlParam
映射查询字符串或表单参数
*/
func UrlParam() {
	router := gin.Default()
	router.POST("/postParam", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		log.Printf("ids: %v， names: %v", ids, names)
	})
	router.Run()
}

/*
QueryParam
查询字符串参数
*/
func QueryParam() {
	router := gin.Default()
	// 使用现有的基础请求对象解析查询字符串参数
	// 示例url: /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname", "Guanyu")
		lastName := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	})
	router.Run()
}
