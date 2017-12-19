package main

import (
	"net/rpc"
	"log"
	"net/http"
	"learntogo/gopl/ch8/rpc1"
	"net"
)


func main() {
	arith:=new(rpc1.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l,e:=net.Listen("tcp",":8001")
	if e!=nil {
		log.Fatalln(e)
	}
	http.Serve(l,nil)
}
