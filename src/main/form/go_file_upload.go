package form

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

/**
文件上传
要使表单能够上传文件，第一步就是要添加form的enctype属性，enctype属性有以下三种情况
	application/x-www-form-urlencoded   表示在发送前编码所有字符（默认）；
	multipart/form-data   不对字符编码。在使用包含文件上传控件的表单时，必须使用该值；
	text/plain    空格转换为 "+" 加号，但不对特殊字符编码。
*/

// UploadFile 增加一个handler
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// 获取请求方法
	fmt.Println("method :", r.Method)
	if r.Method == "GET" {
		// 返回页面
	} else {
		// 移位操作
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Fprintln(w, "%v", handler.Header)
		// 相对路径是以项目根目录为起点
		filePath := "./file/"
		path := filepath.Join(filePath + handler.Filename)
		os.Mkdir(filePath, os.ModeDir)
		openFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		io.Copy(openFile, file)
	}
}
