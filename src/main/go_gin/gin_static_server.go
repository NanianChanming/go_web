package go_gin

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

/*
静态文件服务
*/

func StaticFileServer() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.Run()
}

func AssetStatic() {
	router := gin.New()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(t)
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "/html/index.tmpl", nil)
	})
	router.Run()
}

/*
案例由go-assets-builder实现，此处暂时略过
*/
func loadTemplate() (*template.Template, error) {
	/*t := template.New("")
	for name, file := range {

	}*/
	return nil, nil
}
