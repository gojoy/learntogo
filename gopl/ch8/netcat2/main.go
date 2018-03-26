package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
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
	}()

	inputNum, err := io.Copy(conn, os.Stdin)
	fmt.Println("inputnum is ", inputNum)
	if err != nil {
		log.Fatal(err)
	}
	if inputNum == 0 {
		fmt.Println("input nil!")
	}

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println("copy error", err)
		return
	}
}
