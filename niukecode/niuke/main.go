//practice coding
package main

import (
	"fmt"
	"time"
	"niukecode/huawei"
)

//this is main
func main() {
	t1:=time.Now()

	huawei.CodeLoger()

	fmt.Println(time.Now().Sub(t1))

}


