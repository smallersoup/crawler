package rpcsupport

import (
	"net/rpc"
	"net"
	"net/rpc/jsonrpc"
	"fmt"
	"log"
)

func ServeRpc(host string, service interface{}) error {

	rpc.Register(service)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", host))

	if err != nil {
		return err
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

func NewClient(host string) (*rpc.Client, error) {

	conn, err := net.Dial("tcp", fmt.Sprintf(":%s", host))

	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
