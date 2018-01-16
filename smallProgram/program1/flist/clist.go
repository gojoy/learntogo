package flist

import (
	"sync"
	"errors"
)

var (
	SafeEmptyError=errors.New("Pop Error:Null List\n")
	SafeNotExistError=errors.New("Remove Error:Not Exist\n")
)

type SafeJobList struct {
	sync.Mutex
	Data []string
}

func NewSafeList() *SafeJobList  {

	var (
		w sync.Mutex
		d=make([]string,0)
	)

	return &SafeJobList{
		Data:d,
		Mutex:w,
	}
}

func (l *SafeJobList) Append(v string)  {
	l.Lock()
	defer l.Unlock()
	l.Data=append(l.Data,v)

}

func (l *SafeJobList) Pop() (string, error) {
	if len(l.Data)==0 {
		return "",SafeEmptyError
	}
	l.Lock()
	defer l.Unlock()
	r:=l.Data[0]
	l.Data=l.Data[1:]
	return r,nil
}

func (l *SafeJobList) Remove(v string) error {

	if len(l.Data)==0 {
		return EmptyError
	}

	l.Lock()
	defer l.Unlock()

	for i,j:=range l.Data {
		if j==v {
			l.Data=append(l.Data[:i],l.Data[i+1:]...)
			return nil
		}
	}
	return SafeNotExistError
}
