package xunlei

import (
	"fmt"
	"strconv"
)

func Q1() {
	var (
		n int
	)
	fmt.Scanln(&n)
	fmt.Println(q1(n))
}

func q1(n int) int {
	var (
		flag        bool = false
		num, i, ren int
	)
	if n == 0 {
		return 0
	}
	renum1 := make([]byte, 0)
	if n < 0 {
		flag = true
		num = -n
	} else {
		num = n
	}
	s := strconv.Itoa(num)
	for i = len(s) - 1; i >= 0; i-- {
		if s[i] != '0' {
			break
		}
	}
	for ; i >= 0; i-- {
		renum1 = append(renum1, s[i])
	}

	ren, _ = strconv.Atoi(string(renum1))
	if flag {
		return -ren
	}
	return ren

}
