package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	defer func() {
		fmt.Println("close conn")
		conn.Close()
	}()
	if err != nil {
		log.Fatal(err)
	}

	//if _,err=io.Copy(os.Stdout,conn);err!=nil {
	//	log.Fatal(err)
	//}
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
		return
	}
}
