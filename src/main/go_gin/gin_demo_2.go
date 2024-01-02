package go_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
