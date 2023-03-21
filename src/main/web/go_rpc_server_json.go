package web

import (
	"errors"
)

/*
JSON RPC是数据编码采用了json，而不是gob编码，其他和RPC概念一样。
下面用Go提供的json-rpc标准来实现
*/

type Arith3 int

func (t *Arith3) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith3) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

/*func init() {
	arith3 := new(Arith3)
	rpc.Register(arith3)
	log.Println("--json rpc 服务注册--")

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8081")
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
			fmt.Println("Fatal error ", err.Error())
			return
		}
		jsonrpc.ServeConn(conn)
	}
}*/
