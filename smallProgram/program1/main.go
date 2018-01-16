package main

import (
	"fmt"
	"os"
	"learntogo/smallProgram/program1/fqueue"
	"time"
	"strconv"

)

var (
	i int
	err error
	v string

	done =make(chan error,0)
)

func main() {

	fmt.Printf("argnum %d,args %v\n ",len(os.Args),os.Args)

	jobq:=fqueue.NewfsQ(10)

	go func() {
		for i=0;i<100;i++{
			s:=strconv.Itoa(i)
			jobq.Push(s)
			time.Sleep(200*time.Millisecond)
		}
		jobq.Close()
	}()

	time.Sleep(4000*time.Millisecond)

	go func() {

		v,err=jobq.Pop()
		for err==nil{
			fmt.Printf("%s\t",v)

			v,err=jobq.Pop()
		}
		done<-err
	}()

	fmt.Println("wait")
	fmt.Println(<-done)
	//time.Sleep(2*time.Second)
}

