package main

import (
	"path/filepath"
	"net"
	"log"
	"os"
	"io"
	"fmt"
)

var (
	glog=log.New(os.Stderr,"unixsocket: ",log.Lshortfile)
)

func main() {
	var (
		err error
		sname="my.sock"
		dir="/opt/gofile/bin/tmp"
	)
	fsname:=filepath.Join(dir,sname)
	unixAddr,err:=net.ResolveUnixAddr("unix",fsname)
	if err!=nil {
		glog.Fatalln(err)
	}
	conn,err:=net.DialUnix("unix",nil,unixAddr)
	if err!=nil {
		glog.Fatalln(err)
	}

	_,err=conn.Write([]byte("send by client!!"))
	if err!=nil {
		glog.Printf("write err:%v\n",err)
	}
	fmt.Printf("localaddr is: %v\nremotaddr is: %v\n",
		conn.LocalAddr(),conn.RemoteAddr())
	mustCopy(os.Stdout,conn)
}

func mustCopy(dst io.Writer,src io.Reader)  {
	if _,err:=io.Copy(dst,src);err!=nil {
		glog.Println(err)
		return
	}
}
