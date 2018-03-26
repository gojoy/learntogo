package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
)

var (
	glog = log.New(os.Stderr, "unixsocket: ", log.Lshortfile)
)

func main() {
	var (
		err   error
		mydir = "/opt/gofile/bin/tmp"
	)
	err = os.Chdir(mydir)
	if err != nil {
		log.Fatalln(err)
	}
	sname := filepath.Join(mydir, "my.sock")
	fmt.Printf("sname is %v\n", sname)
	if _, err = os.Stat(sname); err == nil {
		os.Remove(sname)
		glog.Printf("exit socket file:%v\n", sname)
	}

	unixaddr, err := net.ResolveUnixAddr("unix", sname)
	if err != nil {
		glog.Fatalln(err)
	}

	l, err := net.ListenUnix("unix", unixaddr)
	if err != nil {
		glog.Fatalln(err)
	}
	defer l.Close()

	for {

		c, err := l.AcceptUnix()
		if err != nil {
			glog.Fatalln(err)
		}

		go handUnixConn(c)
	}

}

func handUnixConn(c *net.UnixConn) {

	fmt.Printf("laddr: %v\nraddr %v\n", c.LocalAddr(), c.RemoteAddr())

	go func() {

		fmt.Printf("recv from client is")
		io.Copy(os.Stdout, c)
	}()

	defer c.Close()
	for i := 0; i < 10; i++ {
		_, err := io.WriteString(c, "simpleUnixTimeServer:"+time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
