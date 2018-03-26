package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	//fname:="/tmp/testoverlay/merged/c.txt"
	fname := flag.String("file", "", "open file")
	flag.Parse()

	rfp, err := os.Open(*fname)
	if err != nil {
		fmt.Printf("open readonly err:%v", err)
	}
	log.Println("open read only")
	time.Sleep(10 * time.Second)
	rfp.Close()

	log.Println("open write")
	rdfp, err := os.OpenFile(*fname, os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("open rw err:%v", err)
		return
	}
	bs, err := ioutil.ReadAll(rdfp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("context:", string(bs))
	time.Sleep(10 * time.Second)
	rdfp.Write([]byte("add by file\n"))
	time.Sleep(10 * time.Second)
	rdfp.Close()

	return
}
