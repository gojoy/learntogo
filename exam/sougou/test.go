package sougou

import (
	"fmt"
	"sort"
)

type pos [2]int

type set []pos

func Testsort() {
	sets:=make(set,0)
	for i:=0;i<50;i++ {
		tmp:=pos{i,i%7}
		sets=append(sets,tmp)
	}
	//sets=append(sets,pos{1,2})
	fmt.Printf("set is %v\n",sets)
	//sort.Slice(sets, func(i,j int) bool {return sets[i][1]<sets[j][1]})
	sort.Sort(sets)
	fmt.Printf("set is %v\n",sets)
}

func (p set)Len() int  {
	return len(p)
}

func (p set)Less(i, j int) bool {
	return p[i][1]<p[j][1]
}

func (p set)Swap(i, j int) {
	p[i],p[j]=p[j],p[i]
	return
}