package main

import (
	"net"
	"fmt"
	"io"
	"os"
)

func main() {

	listener,err:=net.Listen("tcp","localhost:30000")
	if err!=nil {
		fmt.Println("error listen",err)
		return
	}
	for {
		conn,err:=listener.Accept()
		if err!=nil {
			fmt.Println("error accept",err)
			continue
		}
		go doServe(conn)
	}

}

func doServe(conn net.Conn)  {
	defer conn.Close()
	for {
		if _,err:=io.Copy(os.Stdout,conn);err!=nil{
			break
		}
	}
}
