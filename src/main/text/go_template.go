package text

import (
	"net/http"
	"os"
	"text/template"
)

func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	RangeData()
}

/*
TemplateDemo
go 模板处理
web应用反馈给客户端的信息大部分内容是静态的，采用模板可以复用很多静态代码
在go语言中，可以使用template包进行模板处理，
使用类似Parse、ParseFile、Execute等方法从文件或者字符串加载模板。
*/
func TemplateDemo() {
	t := template.New("first template")
	t.ParseFiles("./file/welcome.html")
	// 取出需要渲染的数据传回给页面渲染
	//t.Execute(w, nil)
}

type Person struct {
	UserName string
}

/*
InsertDataTemplate
在模板中插入数据
go语言的模板通过{{}}来包含需要在渲染时被替换的字段，
{{.}}表示当前对象，如果要访问当前对象的字段通过{{.FieldName}}
需要注意的是，这个字段首字母必须是大写的（相当于java中public修饰）
*/
func InsertDataTemplate() {
	t := template.New("fieldname example")
	t.Parse("hello {{.UserName}}!")
	person := Person{UserName: "Zhangsan"}
	t.Execute(os.Stdout, person)
}

type Person2 struct {
	Person
	Emails []string
}

/*
RangeData
输出嵌套的字段内容
如果字段里还有对象，那么可以使用{{with ...}}...{{end}}和{{range ...}}{{end}}来进行数据输出
·{{range}}和go语法里的range类似，循环操作数据
·{{with}}操作是指当前对象的值，类似上下文的概念
*/
func RangeData() {
	person := Person{UserName: "zhangsan"}
	person2 := Person2{
		Person: person,
		Emails: []string{
			"111@qq.com", "222@qq.com",
		}}
	t := template.New("example")
	t.Parse(`hello {{.UserName}}!
            {{range .Emails}}
                an email {{.}}
            {{end}}`)
	t.Execute(os.Stdout, person2)
}
