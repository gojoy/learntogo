package wangyi

import (
	"fmt"
	"strconv"
)

//对于一个整数X，定义操作rev(X)为将X按数位翻转过来，并且去除掉前导0。例如:
//如果 X = 123，则rev(X) = 321;
//如果 X = 100，则rev(X) = 1.
//现在给出整数x和y,要求rev(rev(x) + rev(y))为多少？
func Q4()  {
	var (
		x,y int
	)
	fmt.Scanln(&x,&y)
	fmt.Printf("%d\n",rev(rev(x)+rev(y)))
}

func rev(x int) int  {
	var (
		sx string
		res int
		err error
	)
	for x%10==0 {
		x=x/10
	}
	//logg.Printf("x is %d\n",x)
	sx=strconv.Itoa(x)
	sx=revs(sx)
	//logg.Printf("sx is %v\n",sx)
	res,err=strconv.Atoi(sx)
	if err!=nil {
		logg.Println(err)
		return -1
	}
	return res
}

func revs(s string) string {
	l:=len(s)
	if l<2 {
		return s
	}
	b:=[]byte(s)
	for i:=0;i<l/2;i++ {
		b[i],b[l-i-1]=b[l-i-1],b[i]
	}
	//logg.Println(string(b))
	return string(b)
}
