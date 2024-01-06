package go_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
CustomLogger
自定义日志文件
*/
func CustomLogger() {
	router := gin.New()
	// LoggerWithFormatter 中间件会写入日志到 gin.DefaultWriter
	// 默认gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/customLogger", func(context *gin.Context) {
		context.String(http.StatusOK, "success")
	})
	router.Run()
}
