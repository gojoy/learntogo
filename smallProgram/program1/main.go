package main

import "fmt"

const CREATE int = 1
const STOP int  = 2

var rdetach  =false

func main() {
	action:=STOP
	var (
		detach=rdetach || (action==CREATE)
	)
	fmt.Printf("detach is %v\n",detach)
}

