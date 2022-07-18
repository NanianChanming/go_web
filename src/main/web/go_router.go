package web

import (
	"net/http"
)

/**
go支持外部路由实现的路由器，ListenAndServe的第二个参数就是用以配置外部路由器的，
它是一个handler接口，即外部路由只要实现了Handler接口即可
*/
// 实现一个简易路由器

type MyMux struct {
}

func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		HelloWorldServer(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

/**
go处理http请求的代码执行过程
1.首先调用http.HandlerFunc,按顺序做了几件事：
	a.调用了DefaultServeMux的HandlerFunc
	b.调用了DefaultServeMux的Handle
	c.往DefaultServeMux的map[String]muxEntry中增加对应的handler和路由规则
2.其次调用了http.ListenAndServe, 按顺序做了几件事：
	a.实例化Server
	b.调用Server的ListenAndServe()
	c.调用netListen("tcp",addr)监听端口
	d.启动一个for循环，在循环体中Accept请求
	e.对每一个请求实例化一个Conn,并且开启一个goroutine处理
	f.读取每个请求的内容 w, err := c.readRequest()
	g.判断handler是否为空，如果没有设置handler，就设置为DefaultServeMux
	h.调用handler的ServeHttp

*/
