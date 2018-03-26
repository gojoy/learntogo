package jd

import (
	"fmt"
	"math"
)

func Q1() {
	var (
		n int64
		i int64
	)
	fmt.Scanln(&n)
	fab := make([]int64, n)
	fab[1] = 1
	fab[2] = 3
	for i = 3; i < n; i++ {
		fab[i] = fab[i-1] + i
		if fab[i] >= n {
			break
		}
	}
	fmt.Println(i)
}

func Q1version1() {
	var (
		n     uint64
		n1, x float64
	)
	fmt.Scanln(&n)
	if n == 1 {
		fmt.Printf("1\n")
		return
	}
	n1 = float64(n)
	x = math.Sqrt(2*n1+0.25) - 0.5
	//fmt.Printf("x is %v\n",x)
	if x-float64(uint64(x)) > 0 {
		fmt.Printf("%d\n", uint64(x)+1)
		return
	}
	fmt.Println(uint64(x))
}
