package go_project

import (
	"reflect"
	"regexp"
)

/*
http路由组件负责将HTTP请求交到对应的函数处理，或者是一个struct的方法
路由在框架中相当于一个事件处理器，而这个事件包括：
·用户请求的路径（path 例如：/user/123），当然还有查询串信息 例如:?id=11
·http请求的方法（method）（GET、POST等）
路由器就是根据用户请求的事件信息转发的相应的处理函数（控制层）
路由的思想主要集中在两点：
·添加路由信息
·根据用户请求转发到要执行的函数
Go默认的路由添加是通过函数http.handle和http.HandleFunc等来实现的，
底层都是调用了DefaultServeMux.Handle(pattern string, handler Handler),
这个函数会把路由信息存储在一个map信息中map[string]muxEntry,这就解决了上面说的第一点
Go监听端口，然后接收到tcp连接会给Handler来处理，上面的例子默认nil即为http.DefaultServeMux,
通过DefaultServeMux.ServeHTTP函数来进行调度，遍历之前存储的map路由信息，和用户访问的URL进行匹配，
以查询对应注册的处理函数，这样就实现了上面所说的第二点。

目前几乎所有的应用路由都是基于http默认的路由器，但是Go自带的路由器有几个限制
·不支持参数设定，例如/user/:uid这种泛类型匹配
·无法很好的支持REST模式，无法限制访问的方法，例如上面的例子中，用户访问/foo, 可以用GET、POST、DELETE、HEAD等方式访问
·一般网站的路由规则太多了，编写繁琐，这种路由多了之后其实可以进一步简化，通过struct的方法进行一种简化
beego框架的路由器基于上面的几点限制考虑设计了一种REST方式的路由实现，路由设计也是基于上面Go默认设计的两点来考虑：
存储路由和转发路由
针对前面所说的限制点，我们首先要解决参数支持就需要用到正则，第二和第三点我们通过一种变通的方法解决。
REST的方法对应到struct的方法中去，然后路由到struct而不是函数，这样在转发路由的时候就可以根据method来执行不同的方法
根据上面的思路，我们设计两个数据类型controllerInfo（保存路径和对应的struct，这里是一个reflect.Type类型）
和ControllerRegistor（routers是一个slice用来保存用户添加的路由信息，以及beego框架的应用信息）
*/
type controllerInfo struct {
	regex          *regexp.Regexp
	params         map[int]string
	controllerType reflect.Type
}

type ControllerRegistor struct {
	routers []*controllerInfo
	//Application *App
}
