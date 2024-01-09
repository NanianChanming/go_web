package go_gin

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

/*
运行多个服务
*/

var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "Welcome server 01",
		})
	})
	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(context *gin.Context) {
		context.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})
	return e
}

func MutiServer() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

/*
RedirectURL
重定向
http重定向很容易，内部外部重定向均支持
*/
func RedirectURL() {
	router := gin.Default()
	router.GET("/test", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})
	// 通过post方法进行HTTP重定向。
	router.POST("/test", func(context *gin.Context) {
		context.Redirect(http.StatusFound, "/foo")
	})

	// 路由重定向使用HandleContext
	router.GET("/test01", func(context *gin.Context) {
		context.Request.URL.Path = "/test2"
		router.HandleContext(context)
	})

	router.GET("/test2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	router.Run()
}
