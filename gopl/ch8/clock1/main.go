package main

import (
	"net"
	"log"
	"io"
	"time"
	"os"
)

var (
	glog=log.New(os.Stderr,"clock1",log.Lshortfile)
)

func main() {
	listener,err:=net.Listen("tcp",":8000")
	if err!=nil {
		glog.Fatal(err)
	}
	for {
		conn,err:=listener.Accept()
		if err!=nil {
			glog.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn)  {
	defer c.Close()
	for {
		_,err:=io.WriteString(c,time.Now().Format("15:04:05\n"))
		if err!=nil {
			glog.Print(err)
			return
		}
		time.Sleep(1*time.Second)
	}
}
