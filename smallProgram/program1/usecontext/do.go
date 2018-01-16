package main

import (
	"context"
	"fmt"
	"time"
	"os"
	"sync"
)

func main() {
	doHandle()
}

func doHandle() {
	wg:=&sync.WaitGroup{}
	//_, cancel := context.WithCancel(context.TODO())
	wg.Add(1)
	go doStuff(context.Background(),wg)
	time.Sleep(10 * time.Second)
	//cancel()
	fmt.Println("cancle!")
	time.Sleep(1*time.Second)
	wg.Wait()
	return
}

func doStuff(ctx context.Context,wg *sync.WaitGroup) {
	defer wg.Done()
	for i:=0;i<15;i++{
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stderr,"done!\n")
			return
		default:
			fmt.Println("work")
		}
	}
}
