package main

import (
	"log"
	"os"
	"net"
	"fmt"
)

var (
	glog=log.New(os.Stderr,"udpserver: ",log.Lshortfile)
)

func main() {
	buf:=make([]byte,512)
	updAddr,err:=net.ResolveUDPAddr("udp","localhost:8001")
	if err!=nil {
		glog.Fatalln(err)
	}

	c,err:=net.DialUDP("udp",nil,updAddr)
	if err!=nil {
		glog.Println(err)
	}

	fmt.Printf("localaddr is: %v\nremotaddr is: %v\n",
		c.LocalAddr(),c.RemoteAddr())

	_,err=c.Write([]byte("hi"))
	//_,err=c.WriteTo([]byte("hi"),c.RemoteAddr())
	if err!=nil {
		glog.Fatalln(err)
	}

	n,_,err:=c.ReadFromUDP(buf[0:])
	if err!=nil {
		glog.Println(err)
	}
	fmt.Printf("recv:%v\n",string(buf[:n]))
}


