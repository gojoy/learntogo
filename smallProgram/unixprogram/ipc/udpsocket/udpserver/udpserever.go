package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var (
	glog  = log.New(os.Stderr, "udpserver: ", log.Lshortfile)
	count uint
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8001")
	if err != nil {
		glog.Fatalln(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		glog.Fatalln(err)
	}

	for {
		handleConn(conn)
	}

}

func handleConn(c *net.UDPConn) {

	buf := make([]byte, 512)
	count++
	fmt.Printf("localaddress is %v\n", c.LocalAddr())

	n, reveaddr, err := c.ReadFromUDP(buf[0:])
	if err != nil {
		glog.Println(err)
	}
	fmt.Printf("recv:%v\nremote address: %v", string(buf[:n]), reveaddr)
	_, err = c.WriteToUDP(
		[]byte(fmt.Sprintf("send by ggudp server,count %v\n", count)),
		reveaddr)
	if err != nil {
		glog.Println(err)
	}

}
