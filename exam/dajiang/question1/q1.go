package main

import (
	"fmt"
)

func main() {
	Q1()
	//fmt.Printf("%d\n",9^1)
}

func Q1()  {
	var (
		n,num int
		inputnums int

	)
	in,err:=fmt.Scanf("%d",&n)
	if err!=nil || in!=1{
		fmt.Scanf("%d",&n)
	}
	//fmt.Printf("%v.%v",in,err)
	//if in!=1 || err!=nil {
	//	fmt.Scanf("%d",&n)
	//}
	//fmt.Printf("n is %d\n",n)
	nums:=make([]int,0)

	for {
		if inputnums>=n {
			break
		}else {
			in,err:=fmt.Scanf("%d",&num)
			if in==1 && err==nil {
				nums=append(nums,num)
				inputnums++
			}
		}
	}
	//fmt.Printf("nums is %d\n",nums)
	fmt.Printf("%d\n",counts(nums))
}


func counts(n []int)int  {
	var (
		res int
	)
	l:=len(n)
	if l<2 {
		return 0
	}
	for i:=0;i<l-1;i++ {
		for j:=i+1;j<l;j++ {
			res=res+countnum(n[i],n[j])
		}
	}
	return res
}


func countnum(a,b int) int  {
	var (
		res int
		lasta,lastb int
	)
	for a>0 || b>0 {
		//fmt.Printf("a,b %b,%b\n",a,b)
		lasta=a&1
		lastb=b&1
		res=res+(lasta^lastb)
		//fmt.Printf("la is %d,lb is %d,res is %d\n",lasta,lastb,lasta^lastb)
		a=a>>1
		b=b>>1
	}
	return res
}