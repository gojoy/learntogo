package main

import (
	"net"
	"log"
	"io"
	"os"
	"fmt"
)

func main() {
	conn,err:=net.Dial("tcp","localhost:8000")
	if err!=nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Println("close conn")
		if err:=conn.Close();err!=nil {
			log.Println("close err ",err)
		}
	}()
	//if _,err=io.Copy(os.Stdout,conn);err!=nil {
	//	log.Fatal(err)
	//}
	mustCopy(os.Stdout,conn)
}

func mustCopy(dst io.Writer,src io.Reader)  {
	if _,err:=io.Copy(dst,src);err!=nil {
		log.Println(err)
		return
	}
}