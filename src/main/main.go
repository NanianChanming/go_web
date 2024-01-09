package main

import (
	_ "go_web/src/main/deploy"
	_ "go_web/src/main/go_excel"
	"go_web/src/main/go_gin"
)

func main() {
	go_gin.MutiServer()
}
