package main

import (
	"errors"
	"fmt"
	"goplay/rpcp/jsonrpcp/common"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

// Arith .
type Arith int

// Multiply .
func (t *Arith) Multiply(args *common.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide .
func (t *Arith) Divide(args *common.Args, quo *common.Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
