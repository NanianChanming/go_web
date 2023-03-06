package http

// go 网络编程
/**
go是如何接收客户端请求的
在执行 http.ListenAndServe监控端口之后，底层调用了func (srv *Server) Serve(l net.Listener)方法，
这个函数就是处理接收客户端的请求信息。这个函数里起了一个for{}循环，首先通过Listener接收请求，其次创建一个conn。
最后单独开了一个goroutine， go c.serve() 把这个请求的数据当作参数扔给这个conn去服务。
这个就是高并发体现了，用户的每一次请求都是一个新的goroutine去服务，相互不影响。

如何具体分配到相应的函数来处理请求
conn首先会解析request c.readRequest(), 然后获取相应的handler handler := c.server.Handler
就是我们刚才在调用函数ListenAndServe 时候的第二个参数。这个变量就是一个路由器，用来匹配url跳转到其相应的handle函数。
这个设置是我们在调用代码的第一句 http.HandleFunc("/", xxx) 的时候注册了请求路由规则。
*/
