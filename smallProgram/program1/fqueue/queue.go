package fqueue

import (
	"errors"
)

type fsQ chan string

var (
	EndQueue =errors.New("queue End\n")
)

func NewfsQ(n uint) fsQ {
	c:=make(fsQ,n)
	return c
}

func (f *fsQ) Push(v string) error {
	*f<-v
	return nil
}

func (f *fsQ) Pop() (string,error) {


	if v,ok:=<-*f;ok {
		return v,nil
	}
	return "",EndQueue
}

func(f *fsQ) Close() {
	close(*f)
}
