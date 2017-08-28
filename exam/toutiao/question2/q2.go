package main

import "fmt"

func main() {
	var (
		n int
		err error
	)
	_,err=fmt.Scanf("%d",&n)
	if err!=nil {
		return
	}

	arr:=make([]int,n)
	for i:=0;i<n;i++ {
		fmt.Scanf("%d",&arr[i])
	}
	if n==0 {
		fmt.Printf("0")
		return
	}
	fmt.Printf("%d\n",getmax(arr))

}

func getmax(a []int) int {
	var (
		l int=len(a)
		currvalue,maxvalue,currsum,currmin int
	)
	if l==1 {
		return a[0]*a[0]
	}
	maxvalue=a[0]*a[0]
	currmin=a[0]
	for i:=0;i<l;i++ {
		valuenext:=min(currmin,a[i])*(currsum+a[i])
		nowvalue:=a[i]*a[i]
		if nowvalue>valuenext {
			currvalue=nowvalue
			currmin=a[i]
			currsum=a[i]
		} else {
			currvalue=valuenext
			currsum=currsum+a[i]
			currmin=min(currmin,a[i])
		}
		maxvalue=max(maxvalue,currvalue)
	}
	return maxvalue
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}