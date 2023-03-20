package unit_test

import "net/http"

/*
Go语言的设计准则是：简洁、明白，
简洁是指语法和C类似，相当的简单，
明白是指任何语句都是很明显的，不含有任何隐含的东西，在错误处理方面的设计中也贯彻了这一思想。
在C语言里面是通过返回-1和NULL之类的信息来表示错误，但是对于使用者来说，
不查看相应的API说明文档，根本搞不清楚这个返回值究竟代表什么，比如返回0是成功还是失败
而Go定义了一个叫做error的类型，来显式表达错误，
在使用时，通过把返回的error变量与nil的比较，来判断操作是否成功。
例如：os.Open函数在打开文件失败时将返回一个不为nil的error变量
func Open(name string)(file *File, err error)
*/
/*
Error 类型
error类型是一个接口类型，定义如下：
type error interface{
	Error() string
}
error 是一个内置的接口类型，我们可以在/builtin/包下面找到相应的定义。
而我们在很多内部包里面用到的error是errors包下面的实现的私有结构errorString
*/
func init() {
	http.Handle("/view", appHandler(viewRecord))
}

func viewRecord(w http.ResponseWriter, r *http.Request) *appError {
	// 返回详细信息
	/*if err := datastore.Get(c, key, record); err != nil {
		return &appError{err, "Record not found", 404}
	}
	if err := viewTemplate.Execute(w, record); err != nil {
		return &appError{err, "Can't display record", 500}
	}*/
	return nil
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//if err := fn(w, r); err != nil {
	//	http.Error(w, err.Error(), 500)
	//}
	// 自定义路由器改造
	if e := fn(w, r); e != nil { // e is *appError, not os.Error
		http.Error(w, e.Message, e.Code)
	}
}

// 自定义错误信息
type appError struct {
	Error   error
	Message string
	Code    int
}

/*
如上所示，在我们访问view的时候可以根据不同的情况获取不同的错误码和错误信息，
虽然这个和第一个版本代码量差不多，但这个显示的错误更加明显，提示的错误信息更加友好，
扩展性也比第一个好
*/
