package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Printf("q3\n")
	v1()
}

func v1() {
	scan:=bufio.NewScanner(os.Stdin)
	output:=bufio.NewWriter(os.Stderr)
	for scan.Scan() {
		t:=scan.Text()
		t=t+"add by fmt"
		fmt.Println(t)
		output.WriteString(t)
		output.Flush()
	}
	//fmt.Println(scan.Err().Error())
	output.Flush()
}
