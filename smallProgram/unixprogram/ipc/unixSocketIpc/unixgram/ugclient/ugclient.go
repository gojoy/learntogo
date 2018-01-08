package main

import (
	"log"
	"os"
	"path/filepath"
	"net"
	"fmt"
	"strconv"
)

var (
	glog=log.New(os.Stderr,"unixsocket: ",log.Lshortfile)
)

func main() {

	var (
		err error
		sname="myg.sock"
		dir="/opt/gofile/bin/tmp"
	)
	fname:=filepath.Join(dir,sname)

	rAddr,err:=net.ResolveUnixAddr("unixgram",fname)
	if err!=nil {
		glog.Fatalln(err)
	}
	pidname:=strconv.Itoa(os.Getpid())
	clname:=filepath.Join(dir,"gg"+pidname+".sock")
	glog.Printf("clname is %v\n",clname)

	lAddr,err:=net.ResolveUnixAddr("unixgram",clname)
	if err!=nil {
		glog.Fatalln(err)
	}

	//clientListener,err:=net.ListenUnixgram("unixgram",lAddr)
	//defer clientListener.Close()
	//if err!=nil {
	//	glog.Fatalln(err)
	//}

	c,err:=net.DialUnix("unixgram",lAddr,rAddr)
	defer os.Remove(clname)
	if err!=nil {
		glog.Printf("dail err:%v\n",err)
		return
	}

	handUnixgram(c)

}

func handUnixgram(c *net.UnixConn)  {
	defer c.Close()

	buf:=make([]byte,512)

	fmt.Printf("localaddr is: %v\nremotaddr is: %v\n",
		c.LocalAddr(),c.RemoteAddr())

	_,err:=c.Write([]byte("hi!"))
	if err!=nil {
		glog.Println(err)
		return
	}

	glog.Println("begin to read")
	n,err:=c.Read(buf[0:])
	if err!=nil {
		glog.Println(err)
		return
	}
	fmt.Printf("recv from server:%v\n",string(buf[:n]))
}