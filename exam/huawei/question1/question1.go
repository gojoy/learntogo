package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		a   int
		b   int
		err error
	)
	_, err = fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		fmt.Printf("-1")
	}

	x := strconv.Itoa(a)
	y := strconv.Itoa(b)
	fmt.Println(findMatch(x, y))
}

func findMatch(a, b string) int {
	if len(a) != len(b) {
		return -1
	}
	var (
		res int
		max int
	)
	for i := 0; i < len(a); i++ {
		max = 0
		for j := i; j < len(a); j++ {
			if a[j] == b[j] {
				max++
			} else {
				break
			}
		}
		if res < max {
			res = max
		}
	}
	if res <= 1 {
		return 0
	}
	return res
}
