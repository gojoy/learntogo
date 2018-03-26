package main

import (
	"learntogo/gopl/ch8/rpc2"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	l, e := net.Listen("tcp", "localhost:8002")
	if e != nil {
		log.Fatalln("listen error: ", e)
	}
	defer l.Close()

	srv := rpc.NewServer()
	arith := new(rpc2.Arith)
	if e = srv.Register(arith); e != nil {
		log.Fatalln("reg error: ", e)
	}
	for {
		conn, e := l.Accept()
		if e != nil {
			log.Fatalln("accept err: ", e)
		}
		go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
