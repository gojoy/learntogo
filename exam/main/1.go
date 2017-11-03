package main

import (

	"fmt"
	"time"
	_"learntogo/exam/aiyiqi"
	"learntogo/exam/pearls"
	"learntogo/exam/yy"
)

func main() {
	start:=time.Now()
	//fmt.Println(pearls.RandIntByMap(10,200))
	fmt.Printf("begin \n")
	s:=make([]float32,4)
	s[0]=0
	s[1]=0.3
	s[2]=0.2
	s[3]=0.5
	 for i:=0;i<100;i++ {
		go fmt.Println(pearls.Roll(s))
	}
	yy.Q1()
	fmt.Println(time.Since(start))
}
