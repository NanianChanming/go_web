package main

import (
	_ "go_web/src/main/deploy"
	"go_web/src/main/gin"
	_ "go_web/src/main/go_excel"
)

func main() {
	gin.InitGin()
}
