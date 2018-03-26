package main

import (
	"fmt"
	"learntogo/exam/didi/mianshi"
	"time"
)

func main() {
	start := time.Now()
	//fmt.Println(pearls.RandIntByMap(10,200))
	mianshi.Tqsort()
	fmt.Println(time.Since(start))
}
