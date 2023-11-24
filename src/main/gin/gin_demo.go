package gin

import "github.com/gin-gonic/gin"

/**
gin 快速入门
1.安装gin:
go get -u github.com/gin-gonic/gin
*/

func InitGin() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "success"})
	})
	r.Run()
}
