package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(2)
}

func main() {
	var x, y int
	go func() {
		x = 1
		fmt.Printf("y:%v\t", y)
	}()

	go func() {
		y = 1
		fmt.Printf("x:%v\t", x)
	}()
	time.Sleep(100 * time.Millisecond)
}
