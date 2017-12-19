package main

import (
	"net/rpc"
	"log"
	"learntogo/gopl/ch8/rpc1"
	"fmt"
)

func main() {
	var (
		reply int
	)
	client,err:=rpc.DialHTTP("tcp","localhost:8001")
	if err!=nil {
		log.Fatal(err)
	}
	args:=rpc1.Args{16,4}
	err=client.Call("Arith.Multiply",args,&reply)
	if err!=nil {
		log.Fatalln("rpc call err ",err)
	}
	fmt.Printf("rpc result is %v\n",reply)

}
