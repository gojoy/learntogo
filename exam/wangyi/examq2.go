package wangyi

import "fmt"

func EXAMQ2() {
	var (
		s                string
		i, count, length int
		old              byte
	)
	fmt.Scanln(&s)
	length = len(s)
	if length == 1 {
		fmt.Printf("1.00\n")
		return
	}
	count = 1
	old = s[0]
	for i = 1; i < length; i++ {
		if s[i] == old {
			continue
		} else {
			count++
			old = s[i]
		}
	}
	//fmt.Printf("count is %v\n",count)
	fmt.Printf("%.2f\n", float32(length)/float32(count))
}
