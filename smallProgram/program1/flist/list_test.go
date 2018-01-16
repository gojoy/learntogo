package flist

import (
	"testing"
	"strconv"
	"fmt"
)

func tTestN(t *testing.T) {

	j:=NewJobList()
	for i:=0;i<100;i++ {
		v:=strconv.Itoa(i)
		j.Append(v)
	}

	for i:=0;i<10;i++ {
		v,err:=j.Pop()
		if err!=nil {
			println(err)
			t.FailNow()
			return 
		}
		fmt.Println(v)
	}

	fmt.Printf("len is %v\n",len(j.Data))
	fmt.Println("remove 50")

	if err:=j.Remove("50");err!=nil {

		fmt.Println(err)
		t.FailNow()
		return

	}

	fmt.Printf("len is %v\n",len(j.Data))

	v,err:=j.Pop()
	for err!=EmptyError {
		fmt.Println(v)
		v,err=j.Pop()
	}
	fmt.Printf("len is %v\n",len(j.Data))
	fmt.Printf("for err:%v\n",err)

}

func BenchmarkNewJobList(b *testing.B) {

	l:=NewJobList()
	for i:=0;i<b.N;i++ {
		l.Append(strconv.Itoa(i))

	}
}
