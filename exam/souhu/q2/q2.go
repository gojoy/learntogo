package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {

	input:=bufio.NewReader(os.Stdin)
	s,err:=input.ReadString('a')
	for err==nil {
		fmt.Printf("input is %s\n",s)
		fmt.Printf("bufferd is %d\n",input.Buffered())
		s,err=input.ReadString('a')
	}
	fmt.Println(err)
}
