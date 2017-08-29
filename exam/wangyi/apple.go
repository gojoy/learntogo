package wangyi

import "fmt"

func Q6()  {
	var (
		n int
	)
	fmt.Scanln(&n)
	fmt.Println(buy(n))
}

func buy(n int) int  {
	if n%2==1 {
		return -1
	}
	if n%8==0 {
		return n/8
	}
	return n/8+1
}