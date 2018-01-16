package flist

import (
	"testing"
	"strconv"
	"fmt"
	"math/rand"
	"time"
)

func BenchmarkSafeJobList_Append(b *testing.B) {
	l:=NewSafeList()

	for i:=0;i<b.N;i++ {
		l.Append(strconv.Itoa(i))
	}

}

func TestNewSafeList(t *testing.T) {

	var (
		l=NewSafeList()
		r1=make(chan int)
		r2=make(chan int)
		done=make(chan struct{})

	)

	for i:=0;i<1000;i++ {
		l.Append(strconv.Itoa(i))
	}

	go func() {
		fmt.Printf("begin remove\n")
		var (
			c,i int
		)
		for i=0;i<50;i++ {
			x:=rand.Intn(999)
			if err:=l.Remove(strconv.Itoa(x));err!=nil {
				fmt.Printf("remove %d failed:%v\n",x,err)
				
			}
			c++
			time.Sleep(10*time.Millisecond)
		}
		fmt.Printf("remove %d finish,i is %d\n",c,i)
		done<- struct{}{}
	}()


	go func() {

		var i int
		v,err:=l.Pop()
		for err==nil {
			i++
			time.Sleep(10*time.Millisecond)
			fmt.Printf("%s ",v)
			v,err=l.Pop()
		}

		//for i:=0;i<500;i++ {
		//	if v,err:=l.Pop();err!=nil {
		//		t.Error(err)
		//		t.FailNow()
		//		return
		//	} else {
		//		fmt.Printf("%s ",v)
		//	}
		//	time.Sleep(10*time.Millisecond)
		//}
		r1<-i
		fmt.Println("r1 finish")
	}()

	go func() {

		var i int
		v,err:=l.Pop()
		for err==nil {
			i++
			//time.Sleep(11*time.Millisecond)
			fmt.Printf("%s ",v)
			v,err=l.Pop()
		}

		//for i:=0;i<500;i++ {
		//	time.Sleep(10*time.Millisecond)
		//	if v,err:=l.Pop();err!=nil {
		//		t.Error(err)
		//		t.FailNow()
		//		return
		//	}else {
		//		fmt.Printf("%s ",v)
		//	}
		//
		//}
		r2<-i
		fmt.Println("r2 finish")
	}()



	fmt.Printf("\nr1 %d,r2 %d\n",<-r1,<-r2)
	<-done

	fmt.Println("main end")
}
