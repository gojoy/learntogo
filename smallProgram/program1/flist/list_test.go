package flist

import (
	"testing"
	"strconv"
	"fmt"
)

func TestNewJobList(t *testing.T) {

	j:=NewJobList()
	for i:=0;i<100;i++ {
		v:=strconv.Itoa(i)
		j.Append(v)
	}

	for i:=0;i<10;i++ {
		v,err:=j.Pop()
		if err!=nil {
			println(err)
			return 
		}
		fmt.Println(v)
	}

	if err:=j.Remove("50");err!=nil {

		fmt.Println(err)

	}

	v,err:=j.Pop()
	for err==nil {
		fmt.Println(v)
	}
	t.Error(err)
	t.FailNow()
}
