package main

import (
	"errors"
	"fmt"
	"goplay/rpcp/httprpcp/common"
	"net/http"
	"net/rpc"
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
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
