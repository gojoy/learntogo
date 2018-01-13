package main

import (

	"fmt"
	"time"
	"learntogo/exam/didi/mianshi"
)

func main() {
	start:=time.Now()
	//fmt.Println(pearls.RandIntByMap(10,200))
	mianshi.Tqsort()
	fmt.Println(time.Since(start))
}
