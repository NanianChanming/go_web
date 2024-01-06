package go_gin

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

/*
go 模型绑定和验证
要将请求体绑定到结构体中，使用模型绑定。Gin目前支持JSON、XML、YAML和标准表单值的绑定（foo=bar＆boo=baz）
Gin使用go-playground/validator/v10 进行验证。
使用时，需要在绑定的所有字段上，设置相应的tag, 例如使用JSON绑定时，设置字段标签为`json:"fieldname"`
Gin提供了两类绑定方法
·Type-Must bind
	·Methods -bind,BindJSON,BindXML,BindQuery,BindYAML
	·Behavior -这些方法属于MustBindWith的具体调用。如果发生绑定错误，则请求终止，并触发c.AbortWithError(400, err).SetType(ErrorTypeBind)。
	响应状态码被设置为400并且Content-Type被设置为text/plain;charset=utf-8。
	如果您在此之后尝试设置响应状态码，Gin会输出日志：
	[GIN-debug] [WARNING] Headers were already written. Wanted to overwrite status code 400 with 422.
	如果希望更好控制绑定，考虑使用ShouldBind等效方法。
·Type-Should bind
	·Methods - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML
	·Behavior - 这些方法属于ShouldBindWith的具体调用，如果发生绑定错误，Gin会返回错误并由开发者处理错误和请求。
使用Bind方法时，Gin会尝试根据Content-Type推断如何绑定。如果明确知道要绑定什么，可以使用MustBindWith或ShouldBindWith。
也可以指定必须绑定的字段，如果一个字段tag上加上了binding:"required", 但绑定时是空值，Gin会报错。
*/

type LoginInfo struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func LoginInfoBind() {
	router := gin.Default()
	// 绑定 JSON({"user": "", "password": ""})
	router.POST("/loginJsonBind", func(context *gin.Context) {
		var json LoginInfo
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "manu" || json.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
		context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.POST("/loginFormBind", func(context *gin.Context) {
		var form LoginInfo
		// 根据Content-Type Header 推断使用哪个绑定器
		if err := context.ShouldBind(&form); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "manu" || form.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	router.Run()
}

type Person2 struct {
	Name     string `form:"name"`
	Address  string `form:"address"`
	Birthday string `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func StartPage() {
	router := gin.Default()
	var person Person2
	// 如果是get请求，只使用Form绑定引擎（`query`）
	// 如果是post请求，首先检查`content-type`是否为json或者xml,然后再使用form(form-data)
	router.Any("/bindPerson", func(context *gin.Context) {
		if context.ShouldBind(&person) == nil {
			log.Println(person.Name)
			log.Println(person.Address)
			log.Println(person.Birthday)
		}
		context.String(http.StatusOK, "Success")
	})
	router.Run()
}

/*
绑定表单数据到自定义结构体
*/

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(context *gin.Context) {
	var b StructB
	context.Bind(&b)
	context.JSON(http.StatusOK, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(context *gin.Context) {
	var b StructC
	context.Bind(&b)
	context.JSON(http.StatusOK, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(context *gin.Context) {
	var b StructD
	context.Bind(&b)
	context.JSON(http.StatusOK, gin.H{
		"a": b.NestedAnonyStruct,
		"b": b.FieldD,
	})
}

func GetData() {
	router := gin.Default()
	router.GET("/getb", GetDataB)
	router.GET("/getc", GetDataC)
	router.GET("/getd", GetDataD)
	router.Run()
}

/*
自定义http配置
customHttpConfig
*/
func customHttpConfig() {
	router := gin.Default()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// Logger
func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		// 设置example变量
		context.Set("example", "12345")
		// 请求前
		context.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(latency)
		// 获取发送的status
		status := context.Writer.Status()
		log.Println(status)
	}
}

/*
CustomMiddleware
自定义中间件
*/
func CustomMiddleware() {
	router := gin.New()
	router.Use(Logger())
	router.GET("/custom", func(context *gin.Context) {
		example := context.MustGet("example").(string)
		log.Println(example)
	})
	router.Run()
}
