package main

import (
	"fmt"
	"sort"
)

type pos [3]int

func main() {
	var (
		n int
	)

	fmt.Scanf("%d", &n)

	//arrx:=make([]int,n)
	//arry:=make([]int,n)
	arr := make([]pos, n)
	res := make([]pos, 3)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &arr[i][0], &arr[i][1])
		arr[i][2] = arr[i][0] + arr[i][1]
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i][2] > arr[j][2] })
	for i := 0; i < 3; i++ {
		res[i] = arr[n-i-1]
	}
	sort.Slice(res, func(i, j int) bool { return res[i][0] < res[j][0] })
	for i := 0; i < 3; i++ {
		fmt.Printf("%d %d\n", res[i][0], res[i][1])
	}

	//sort.IntsAreSorted()
}
