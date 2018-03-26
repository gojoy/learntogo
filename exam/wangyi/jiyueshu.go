package wangyi

import "fmt"

func Q5() {
	var (
		x int
	)
	fmt.Scanln(&x)
	fmt.Printf("%d\n", f(x))
}

func f(x int) int {
	if x%2 == 1 {
		return x
	}
	for j := x / 2; j > 0; {
		if x%j == 0 && j%2 == 1 {
			return j
		}
		j = j / 2

	}
	return 1
}

func ff(x int) int {
	var (
		res int
	)
	for i := 1; i <= x; i++ {
		fx := f(i)
		res = res + fx
		//logg.Printf("i is %d,fx is %d,res is %d\n",i,fx,res)
	}
	return res
}
