package main

import (
	"net/rpc/jsonrpc"
	"log"
	"learntogo/gopl/ch8/rpc2"
	"fmt"
)

func main() {
	var res int
	c,e:=jsonrpc.Dial("tcp","localhost:8002")
	if e!=nil {
		log.Fatalln(e)
	}
	args:=rpc2.Args{4,5}
	e=c.Call("Arith.Multiply",args,&res)
	if e!=nil {
		log.Fatalln(e)
	}
	fmt.Printf("rpc call res is %v\n",res)
}
