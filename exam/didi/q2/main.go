package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//最大堆
type zs []int

func main() {
	bufinput()
}

func fmtinput() {
	var (
		k, n, in int
		err      error
	)
	a := make([]int, 0)
	for {
		in, err = fmt.Scanf("%d", &n)
		if in == 0 || err != nil {
			fmt.Println(err)
			break
		} else {
			a = append(a, n)
		}
	}
	fmt.Scanf("%d", &k)
	fmt.Printf("%d\n", q2(a, k))
}

func q2(a []int, k int) int {
	sort.Ints(a)
	//sort.Slice(a, func(i,j int) bool {return a[i]>a[j]})
	return a[len(a)-k]
}

//最小的第k个数
func heapnextknum2(a []int, k int) int {
	var (
		l           int = len(a)
		res, minnum int
		hp          zs = make(zs, k)
		rest        zs = make(zs, l-k)
		r           interface{}
		ok          bool
	)
	copy(hp, a[0:k])
	copy(rest, a[k:])
	fmt.Printf("hp is %v\n rest is %v\n", hp, rest)
	heap.Init(&hp)
	fmt.Printf("after init hp is %v\n", hp)
	for i := 0; i < l-k; i++ {
		minnum = hp[0]

		if rest[i] < minnum {
			heap.Remove(&hp, 0)
			heap.Push(&hp, rest[i])
			fmt.Printf("now remove minnum %d,len is %d\n hp is %v\n", minnum, len(hp), hp)
		}
	}
	r = heap.Pop(&hp)
	res, ok = r.(int)
	if !ok {
		panic("res int error")
	}
	return res

}

//第k个最大的数
func heapnextknum(a []int, k int) int {
	var (
		res  int
		l    int = len(a)
		r    interface{}
		nums zs
		ok   bool
	)
	nums = make(zs, l)
	copy(nums, a)
	fmt.Printf("num t is %T,v is %v\n", nums, nums)
	heap.Init(&nums)
	for i := 0; i < k; i++ {
		r = heap.Remove(&nums, 0)
		fmt.Printf("remove %d,len is %d\n", i, len(nums))
	}

	r = heap.Pop(&nums)
	res, ok = r.(int)
	if !ok {
		panic("res err")
	}
	fmt.Printf("pop is %d\n", res)
	return res

}

func bufinput() {
	var (
		k   int
		err error
	)

	a := make([]int, 0)
	input := bufio.NewReader(os.Stdin)
	b, err := input.ReadString('\n')
	if err != nil {
		fmt.Printf("err is %v\n", err)
		return
	}
	fmt.Printf("b is %v\n", b)
	b = b[:len(b)-1]
	num := strings.Split(b, " ")
	for i := 0; i < len(num); i++ {
		n, err := strconv.Atoi(num[i])
		if err != nil {
			fmt.Printf("atoi err is %v\n", err)
		} else {
			//fmt.Printf("n is %d\n",n)
			a = append(a, n)
		}
	}
	_, err = fmt.Scanf("%d", &k)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("k is %d\n", k)
	fmt.Printf("%d\n", heapnextknum(a, k))
}

func testbufio() error {
	var (
		err error
		b   []byte
		//n int
	)
	//scanner:=bufio.NewScanner(os.Stdin)
	input := bufio.NewReader(os.Stdin)
	b, err = input.ReadBytes('\n')
	for err == nil {
		fmt.Printf("b is %v\n", b)
		b, err = input.ReadBytes('\n')
	}
	fmt.Println(err)
	return err
}
func testscan() {
	var (
		s int
	)
	n, err := fmt.Scan(&s)
	for err == nil {
		fmt.Printf("n is %d,s is %v\n", n, s)
		n, err = fmt.Scanln(&s)
	}
	fmt.Println(err)
}

func (a zs) Less(i, j int) bool {
	return a[i] > a[j]
}
func (a zs) Len() int {
	return len(a)
}
func (a zs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
	return
}
func (a *zs) Pop() interface{} {
	l := len(*a)
	//old:=*a
	res := (*a)[l-1]
	*a = (*a)[:l-1]
	return res
}
func (a *zs) Push(x interface{}) {
	n, ok := x.(int)
	if !ok {
		panic("x err\n")
	}
	*a = append(*a, n)
	return
}
