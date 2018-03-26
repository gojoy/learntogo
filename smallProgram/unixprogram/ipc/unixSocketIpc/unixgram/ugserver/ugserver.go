package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

var (
	glog  = log.New(os.Stderr, "unixsocket: ", log.Lshortfile)
	count uint
)

func main() {
	defer fmt.Println("defer execute!")

	var (
		err   error
		sname = "myg.sock"
		dir   = "/opt/gofile/bin/tmp"
	)
	fname := filepath.Join(dir, sname)
	unixAddr, err := net.ResolveUnixAddr("unixgram", fname)
	if err != nil {
		glog.Fatalln(err)
	}

	conn, err := net.ListenUnixgram("unixgram", unixAddr)
	defer os.Remove(fname)
	if err != nil {
		glog.Fatalln(err)
	}

	defer conn.Close()

	for {
		handUnixgram(conn)
	}

}

func handUnixgram(c *net.UnixConn) {

	buf := make([]byte, 512)

	n, raddr, err := c.ReadFromUnix(buf[0:])
	if err != nil {
		glog.Println(err)
		return
	}
	glog.Printf("laddr: %v\nraddr %v\n", c.LocalAddr(), raddr)
	count++
	fmt.Printf("recv from client:%v\n", string(buf[:n]))

	_, err = c.WriteToUnix(
		[]byte(fmt.Sprintf("send by ggudp server,count %v\n", count)),
		raddr)
	if err != nil {
		glog.Println(err)
	}
}
