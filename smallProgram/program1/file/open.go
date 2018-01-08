package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	//fname:="/tmp/testoverlay/merged/c.txt"
	fname:="/mnt/nfsdir1/c.txt"

	rfp,err:=os.Open(fname)
	defer rfp.Close()
	if err!=nil {
		fmt.Printf("open readonly err:%v",err)
	}

	rdfp,err:=os.OpenFile(fname,os.O_RDWR,0666)
	defer rdfp.Close()
	if err!=nil {
		fmt.Printf("open rw err:%v",err)
		return
	}
	bs,err:=ioutil.ReadAll(rdfp)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("context:",string(bs))

	return
}
