package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type dep [2]uint
type depall []dep

func main() {
	var (
		//a uint
		//b uint
		err error
	)

	inputReader:=bufio.NewReader(os.Stdin)
	str,err:=inputReader.ReadString('\n')
	fmt.Printf("str is %s\n",str)
	if err!=nil {
		sl:=strings.Split(str,",")
		for i:=0;i<len(sl);i++ {
			fmt.Printf("ko\n")
			fmt.Println(sl[i])
		}
		str,err=inputReader.ReadString('\n')
	}
	fmt.Println(err)
}
