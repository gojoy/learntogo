//practice coding
package main

import (
	"fmt"
	"time"
	//"learntogo/niukecode/huawei"

	"learntogo/niukecode/toutiao"
)

//this is main
func main() {
	t1 := time.Now()

	//print("begin\n")
	//huawei.LCS()
	toutiao.Two()

	fmt.Println(time.Now().Sub(t1))
}
