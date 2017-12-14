package list

import (
	"container/list"
	"fmt"
	//"container/heap"
)

type pos [2]int

func TestList() {
	arry:=make([]pos,0)
	for i:=0;i<100;i++ {
		arry=append(arry,pos{i,i+1})
	}
	l:=list.New()
	for i:=0;i<len(arry);i++ {
		l.PushBack(arry[i])
	}
	l.PushBack(pos{1000,-1})
	l.PushFront("begin")

	for e:=l.Front();e!=nil;e=e.Next() {
		fmt.Printf("e is %v\n",e.Value)
	}

}

