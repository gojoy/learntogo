package wangyi

import (
	"fmt"
	"strconv"
)

func EXAMQ1()  {
	var (
		n int
		i int
	)
	num1:=make([]byte,0)
	fmt.Scanln(&n)
	s:=strconv.Itoa(n)

	for i=len(s)-1;i>=0;i-- {
		if s[i]!='0' {
			break
		}
	}
	for ;i>=0;i-- {
		num1=append(num1,s[i])
	}
	nums,_:=strconv.Atoi(string(num1))
	fmt.Println(n+nums)

}
