package web

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

/*
基于tcp协议实现RPC
与http的服务器相比，不同在于：
此处采用了TCP协议，然后需要自己控制连接，当有客户端链接上来后，我们需要把这个连接交给rpc来处理。
如果想要实现多并发，那么可以使用goroutine来实现。
*/

type Arith2 int

func (t *Arith2) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith2) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

/*func init() {
	arith2 := new(Arith2)
	rpc.Register(arith2)
	log.Println("--tcp rpc 服务注册--")

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8090")
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//go tcpRPCServer(conn)
		tcpRPCServer(conn)
	}
}*/

func tcpRPCServer(conn net.Conn) {
	rpc.ServeConn(conn)
	log.Println("--rpc tcp handle--")
}
