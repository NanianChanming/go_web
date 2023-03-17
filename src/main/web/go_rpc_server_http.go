package web

import (
	"errors"
)

/*
Socket和HTTP采用的都是类似“信息交换”模式，即客户端发送一条信息到服务端，然后服务器端会返回一定的信息以表示响应。
客户端和服务端之间约定了交互信息的格式，以便双方都能够解析交互所产生的信息。
但是很多独立的应用并没有采用这种模式，而是采用类似常规的函数调用方式来完成想要的功能。

RPC就是实现函数调用模式的网络化，客户端就像调用本地函数一样，然后客户端就把这些参数打包之后通过网络传递到服务端，
服务端解包到处理过程中执行，然后执行的结果反馈给客户端。

RPC（Remote Procedure Call Protocol）--远程过程调用协议，是一种通过网络从远程计算机程序上请求服务，而不需要了解地城网络技术的协议。
它假定某些传输协议的存在，如TCP或UDP, 以便为通信程序之间携带信息数据，通过它可以使函数调用模式网络化。
在OSI网络通信模型中，RPC跨越了传输层和应用层，RPC使得开发包括网络分布式多程序在内的应用程序更加容易。
*/
/*
Go RPC
go标准包中已经提供了对RPC的支持，而且支持三个级别的RPC：TCP、HTTP、JSONRPC。
但Go的RPC包是独一无二的RPC，它和传统的RPC系统不同，它只支持Go开发的服务器与客户端之间的交互，
因为在内部，它们采用了Gob来编码。

Go RPC的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细要求如下：
·函数必须是导出的（首字母大写）
·必须有两个导出类型的参数
·第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型
·函数还要有一个返回值error
举个例子，正确的RPC函数格式如下：
func (t *T) MethodName(argType T1, replyType *T2) error
T、T1和T2类型必须能被encoding/gob包编解码
任何的RPC都需要通过网络来传递数据，Go RPC可以利用HTTP和TCP来传递数据，
利用HTTP的好处是可以直接复用net/http里面的一些函数
*/

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

/*
通过init函数注册一个Arith的RPC服务，然后通过rpc.HandleHTTP函数把该服务注册到了HTTP协议上，然后我们就可以利用http的方式来传递数据了
*/
//func init() {
//	arith := new(Arith)
//	rpc.Register(arith)
//	rpc.HandleHTTP()
//	http.ListenAndServe(":8080", nil)
//	log.Println("--http rpc 服务注册结束--")
//}
