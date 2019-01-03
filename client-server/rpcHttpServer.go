package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

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

func main() {
	//arith := new(Arith)
	var arith Arith
	rpc.Register(&arith)
	fmt.Println("Register called...")
	rpc.HandleHTTP()
	fmt.Println("HandleHTTP called...")
	err := http.ListenAndServe("127.0.0.1:1234", nil)
	fmt.Println("server started")
	if err != nil {
		fmt.Println(err.Error())
	}
}
