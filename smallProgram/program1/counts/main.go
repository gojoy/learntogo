package main

import (
	"time"
	"fmt"
	"os"
	"log"
)

var (
	startTime=time.Now()
	tick=time.Tick(1*time.Second)
)

func main() {
	dir,err:=os.Getwd()
	if err!=nil {
		log.Fatalln(err)
	}
	fmt.Printf("pwd is %v\n",dir)
	var (
		counts uint
		atob=make(chan int)
		done=make(chan int)
	)
	fmt.Println("start")
	go func() {
		for {
			atob<-1
		}
	}()

	go func() {
		for {
			select {
			case<-atob:
				counts++
			case <-tick:
				fmt.Printf("couns %v\n",counts)
				done<-1
			}
		}
	}()
	<-done
	fmt.Printf("time out,cost %v\n",time.Since(startTime))
}
