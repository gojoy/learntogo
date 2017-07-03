//practice coding
package main

import (
	"fmt"
	"time"
	"learntogo/niukecode/huawei"
)

//this is main
func main() {
	t1:=time.Now()

	huawei.CheckPasswd()

	fmt.Println(time.Now().Sub(t1))

}


