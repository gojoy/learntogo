package main

import "fmt"

func main() {
	var (
		n,m,h int
	)
	a:=make([]int,0)
	_,err:=fmt.Scanln(&n,&m)
	if err!=nil {
		fmt.Println(err)
		return
	}
	for i:=0;i<m;i++ {
		fmt.Scanf("%d",&h)
		//a[i]=h
		a=append(a,h)
	}
	//fmt.Printf("n is %d,m is %d,a is %v\n",n,m,a)
	if m<2 {
		for i:=0;i<n;i++ {
			fmt.Println(a[0])
		}
		return
	}
	getres(a,n)
}

//生成kolakoski 序列
func getres(a []int, n int) {
	var (
		out,in,lengh int
	)
	lengh=0
	res:=make([]int,0)
	for i:=0;i<a[0];i++ {
		res=append(res,a[0])
		lengh++
	}
	out=1
	in=1
	if lengh==1 {
		for ii:=0;ii<a[1];ii++ {
			res=append(res,a[1])

		}
		out++
		if out==len(a) {
			out=0
		}
		in++
	}
	for lengh<=n {
		for i:=0;i<res[in];i++ {
			res=append(res,a[out])
			lengh++
		}
		out++
		in++
		if out==len(a) {
			out=0
		}
	}
	for j:=0;j<n;j++ {
		fmt.Printf("%v ",res[j])
	}
}