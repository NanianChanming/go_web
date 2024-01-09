package go_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
路由参数
*/

func RouterParam() {
	router := gin.Default()
	// 此handler将匹配 /user/john 但不会匹配 /user/ 或者 /user
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello, %s", name)
	})

	// 此handler将匹配 /user/john 和 /user/john/send
	// 如果没有其他路由匹配/user/john, 它将重定向到/user/john/
	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK, message)
	})
	router.Run()
}

/*
RouterGroup
路由组
*/
func RouterGroup() {
	router := gin.Default()
	// 简单的路由组v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// 简单的路由组v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}
	router.Run()
}

func loginEndpoint(context *gin.Context) {
	context.String(http.StatusOK, context.Request.RequestURI)
}

func submitEndpoint(context *gin.Context) {
	context.String(http.StatusOK, context.Request.RequestURI)
}

func readEndpoint(context *gin.Context) {
	context.String(http.StatusOK, context.Request.RequestURI)
}
