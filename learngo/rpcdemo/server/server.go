package main

import (
	"net"
	"net/rpc/jsonrpc"
	"paasadm/pkg/util/log"
	"net/rpc"
	"crawler/learngo/rpcdemo"
)

func main() {

	rpc.Register(rpcdemo.DemoService{})

	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}


