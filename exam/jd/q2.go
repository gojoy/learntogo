package jd

import "fmt"

func Q2()  {
	var (
		l,r int
		count int=0
	)
	l=0
	r=10000
	for i:=l;i<r;i++ {
		if ism(i,0,0) {
			count++
		}
	}
	fmt.Println(count)
}

func ismagicnum(n int) bool  {
	if n<10 {
		return false
	}
	return ism(n,0,0)
}

func ism(n int,sum1 int,sum2 int) bool  {
	if n==0 {
		if sum1==sum2 {
			return true
		}
		return false
	}
	chu:=n/10
	yu:=n%10
	return ism(chu,sum1+yu,sum2) || ism(chu,sum1,sum2+yu)
}
