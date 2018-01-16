package flist

import "errors"

var (
	EmptyError=errors.New("Pop Error:Null List\n")
	NotExistError=errors.New("Remove Error:Not Exist\n")
)

type jobList struct {

	Data []string
}

func NewJobList() *jobList {
	d:=make([]string,0)
	return &jobList{
		Data:d,
	}
}

func (l *jobList) Append(v string) error {
	l.Data=append(l.Data,v)
	return nil
}

func (l *jobList) Pop() (string,error) {
	if len(l.Data)==0 {
		return "",EmptyError
	}
	r:=l.Data[0]
	l.Data=l.Data[1:]
	return r,nil
}

func (l *jobList) Remove(v string) error {
	if len(l.Data)==0 {
		return EmptyError
	}
	for i,j:=range l.Data {
		if j==v {
			l.Data=append(l.Data[:i],l.Data[i+1:]...)
			return nil
		}
	}
	return NotExistError
}
