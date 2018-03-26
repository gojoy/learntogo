package toutiao

import (
	"fmt"
	"sort"
)

//给定整数m以及n各数字A1,A2,..An，将数列A中所有元素两两异或，
// 共能得到n(n-1)/2个结果，请求出这些结果中大于m的有多少个。
func Three() {
	var (
		n, m int
		err  error
	)
	_, err = fmt.Scanf("%d %d", &n, &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	//fmt.Printf("a is %v\n",a)
	deal(a, m, n)
	return
}

func deal(a []int, m, n int) {
	yh := make([]int, 0)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			yh = append(yh, a[i]^a[j])
		}
	}
	//fmt.Printf("yh is %v\n",yh)
	sort.Ints(yh)
	res := sort.Search(len(yh), func(i int) bool { return yh[i] > m })
	fmt.Printf("%d\n", len(yh)-res)

}
