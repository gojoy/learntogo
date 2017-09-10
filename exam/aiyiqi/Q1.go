package aiyiqi

import "fmt"

func EXAMQ1()  {
	var (
		s string
	)
	fmt.Scanln(&s)
	fmt.Printf("count is %v\n",Q1(s))
}

func Q1(s string) int  {
	var (
		count,left,right int
	)
	right=len(s)-1
	for left<right {
		if s[left]=='R' {
			left++
			continue
		}
		if s[right]=='G' {
			right--
			continue
		}
		left++
		right--
		count+=2
	}
	return count

}