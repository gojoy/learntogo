package xunlei

import (
	"fmt"
)

func Q2()  {
	var (
		n,m int
		res int=0
	)
	fmt.Scanln(&n,&m)
	//fmt.Printf("sum is %v,n is %v\n",m,n)
	//fmt.Println(sumofk(m,n))
	sumofkv1(m,n,&res)
	fmt.Printf("%v\n",res)

}


func sumofk(sum int,n int) int  {
	if n<0 || sum<0 {
		return 0
	}
	if sum==n {
		return 1
	}
	return sumofk(sum-n,n-1)+sumofk(sum,n-1)
}

func sumofkv1(sum int,n int,count *int)  {
	if sum<=0 || n<=0 {
		return
	}
	if sum==n {
		*count=*count+1
	}
	sumofkv1(sum-n,n-1,count)
	sumofkv1(sum,n-1,count)
}