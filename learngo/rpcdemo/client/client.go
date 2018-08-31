package main

import (
	"net"
	"net/rpc/jsonrpc"
	"crawler/learngo/rpcdemo"
	"paasadm/pkg/util/log"
)

func main() {

	conn, err := net.Dial("tcp", ":8888")

	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)
	args := rpcdemo.Args{A: 14, B: 3}
	var result float64

	err = client.Call("DemoService.Div", args, &result)

	if err != nil {
		log.Printf("error:%v\n", err)
	} else {
		log.Printf("%d / %d = %.5f\n", args.A, args.B, result)
	}

	args = rpcdemo.Args{A: 14, B: 0}

	err = client.Call("DemoService.Div", args, &result)

	if err != nil {
		log.Printf("error:%v\n", err)
	} else {
		log.Printf("%d / %d = %.5f\n", args.A, args.B, result)
	}
}


