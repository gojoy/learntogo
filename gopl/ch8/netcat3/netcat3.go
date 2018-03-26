package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	done := make(chan int)
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln("dial err", err)
	}
	defer conn.Close()
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Println(err)
		}
		fmt.Println("done")
		done <- 1
	}()
	mustCopy(conn, os.Stdin)
	fmt.Println("receve done is ", <-done)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println("input err:", err)
		return
	}
}
