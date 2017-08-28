package main

import "fmt"

func main() {
	var (
		n,in int
		err error
	)
	a:=make([]int,0)
	in,err=fmt.Scanf("%d",&n)
	for err==nil &&in!=0 {
		a=append(a,n)
		_,err=fmt.Scanf("%d",&n)
	}
	fmt.Printf("%d\n",q1(a))

}

func q1(a []int) int  {
	var (
		currsum int=0
		maxsum int=a[0]
	)
	for j:=0;j<len(a);j++ {
		currsum=max(a[j],currsum+a[j])
		maxsum=max(maxsum,currsum)
	}
	return maxsum
}

func max(a,b int) int  {
	if a > b {
		return  a
	}
	return b
}