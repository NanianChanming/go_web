package web

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

/*
什么是Socket
socket起源于Unix，而unix基本哲学之一就是“一切皆文件”, 都可以用“打开open->读写write/read->关闭close”模式来操作。
Socket就是该模式的一个实现，网络的socket数据传输是一种特殊的I/O，socket也是一种文件描述符。
Socket也具有一个类似于打开文件的函数屌用：Socket(),该函数返回一个整型的Socket描述符，
随后连接的建立、数据传输等操作都是通过该Socket实现的。
常用的Socket类型有两种: 流式Socket(SOCK_STREAM)和数据式Socket(SOCK_DGRAM).
流式是一种面向连接的Socket，针对于面向连接的TCP服务应用;
数据报式Socket是一种无连接的Socket，对应于无连接的UDP服务应用。

Socket如何通信
网络中的进程之间如何通过Socket通信：
首先要解决的问题是如何唯一标识一个进程，否则通信无从谈起，在本地可以通过进程PID来标识一个进程，
但在网络中这是行不通的。TCP/IP协议族已经帮我们解决了这个问题，网络层的“ip地址”可以唯一标识网络中的主机，
而传输层的“协议+端口”可以唯一标识主机中的应用程序（进程）。这样利用三元组（IP地址，协议，端口）就可以标识网络的进程了，
网络中需要互相通信的进程，就可以利用这个标志在他们之间进行交互。
*/

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	//ParseIP()
	serverHandler()
	//clientHandler()
}

func init() {
	//serverHandler()
	//clientHandler()
}

/*
ParseIP
Go支持的IP类型
在Go的net包中定义了很多类型、函数和方法用来网络编程，其中ip的定义如下
type IP []byte
在net包中有很多函数来操作IP，但是其中比较有用的也就几个，
其中ParseIP(s string) IP 函数会把一个IPv4或者IPv6的地址转换为IP类型
*/
func ParseIP() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}

/*
TCP Socket
当我们通过网络端口访问一个服务时，作为客户端来说，我们可以通过向远端某台机器的某个网络端口发送一个请求，
然后得到在机器的此端口上监听的服务反馈信息。
作为服务端，我们需要把服务绑定到某个指定端口，并且在此端口上监听，当有客户端来访问的时候能够读取信息并且写入反馈信息。
在Go语言的net包中有一个类型TCPConn,这个类型可以用来作为客户端和服务端交互的通道，他有两个主要函数：
func (c *TCPConn) Write(b []byte)(int, error)
func (c *TCPComm) Read(b []byte)(int,error)
TCPConn可以在客户端和服务端来读写数据
还有一个TCPAddr类型，他表示一个TCP的地址信息，定义如下
type TCPAddr struct{
	IP IP
	Port int
	Zone string // IPv6 scoped addressing zone
}
在go中通过ResolveTCPAddr获取一个TCPAddr
func ResolveTCPAddr(net, addr string)(*TCPAddr, os.error)
·net参数是“tcp4”、“tcp6”、“tcp”中的任意一个，分别表示TCP(IPv4-only),TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个)
·addr表示域名或者IP地址
*/
/*
TCP client
go中通过net包中的DialTCP函数来建立一个TCP连接，并返回一个TCPConn类型的对象，
当连接建立时服务器端也创建一个同类型的对象，此时客户端和服务器端通过各自拥有的TCPConn对象来进行数据交换。
一般而言，客户端通过TCPConn对象将请求信息发送到服务器端，并返回应答信息，这个连接只有当任一客户端关闭了之后才失效
不然这连接可以一直在使用，建立连接的函数定义如下：
func DialTCP(network string, laddr, raddr *TCPAddr)(*TCPConn, error)
·net参数是“tcp4”，“tcp6”,“tcp”中的任意一个
·laddr表示本机地址，一般设置为nil
·raddr表示远程的服务地址
*/

/*
ClientHandler
一个简单的例子，模拟一个基于http协议的客户端请求去连接一个web服务端。
写一个简单的http请求头，格式类似如下：
"HEAD / HTTP/1.0\r\n\r\n"
*/
func clientHandler() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:post ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

/*
TCP Server
通过net包来创建一个服务器端程序，在服务器端我们需要绑定服务到指定的非激活端口，并监听此端口，
当有客户端请求到达的时候可以接收到来自客户端连接的请求。
net包中有相应功能的函数，定义如下
func ListenTCP(network string, laddr *TCPAddr)(*TCPListener, error)
func (l *TCPListener) Accept()(conn, error)
参数与DialTCP的参数一样。下面实现一个简单的时间同步服务，监听7777端口
*/
func serverHandler() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
