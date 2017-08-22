package toutiao

import (
	"fmt"
	"sort"
)

func Two() {
	var (
		n int
		err error
	)

	_,err=fmt.Scanln(&n)
	if err!=nil {
		fmt.Println(err)
		return
	}
	array:=make([]int,n)
	for i:=0;i<n;i++ {
		fmt.Scanf("%d",&array[i])
	}
 	fmt.Printf("%d\n",dealtwo(array))
}

func dealtwo(a []int) int  {

	var (
		l int=len(a)
		count int=1
		res int

	)
	if l<2 {
		return 3-l
	}
	sort.Ints(a)
	for i:=0;i<l-1;i++ {
		j:=i+1
		if (count%3!=0) {
			if a[j]-a[i]>10 {
				a=append(append(a[:i],a[i]+10),a[j:]...)
				res++
				l++
			}
			i++
			j++
			count++
		}else {
			i++
			j=i+1
		}
	}
	return res
}
