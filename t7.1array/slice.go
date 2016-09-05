package main

import "fmt"

func tmain() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range
}

//使用指针传参对数组赋值
func getpvalue(p1 *[15]int) int {
	for i := 0; i < len(p1); i++ {
		p1[i] = i
	}
	return 1
}

//使用指针传参输出数组
func ptary(p1 *[15]int) int {
	for i := 0; i < len(p1); i++ {
		fmt.Printf("p1[%d] now is %v,and p1 is %v,and addr is %v\n", i, p1[i], p1, &p1[i])
		fmt.Printf("p1 type is %T\n", p1)
	}
	return 1
}

//使用切片传参对数组赋值
func getvalue(p1 []int) {
	for i := 0; i < len(p1); i++ {
		p1[i] = i + 1
	}
}

//使用切片传参输出数组
func outnopoint(p1 []int) {
	for i := 0; i < len(p1); i++ {
		fmt.Printf(" nopoint p[%d] is %v\n", i, p1[i])
	}
}

func byteout(s1 string, num int) {
	var (
		//str string
		s []byte
	)
	//str = "ab cd ef"
	s = []byte(s1)
	fmt.Printf("the string before num is %s\n,and after num is %s,and len is %d\n", s[0:num], s[num:], len(s))
}

func t713(str string) {
	fmt.Printf("result is %s\n", str[len(str)/2:]+str[:len(str)/2])
}
func t714fanz(str string) {
	var (
		s1 []byte
		//s2 []byte
		//n int = len(str)
	)
	s1 = []byte(str)
	//nu := len(s1)
	//s2 := new([n]byte)
	s2 := make([]byte, len(s1))
	for i := 0; i < len(s1); i++ {
		s2[len(s1)-1-i] = s1[i]
	}
	fmt.Printf("exc str is %s\n", s2[:])
}
func mpsort(p1 []int) {
	x := p1[0]
	p2 := make([]int, len(p1))
	p2 = p1
	fmt.Printf("make p2 type is %T,p1[0] is %d,len is %d\n", p2, x, len(p2))
	for i := 0; i < len(p2)-1; i++ {
		for u := 0; u < len(p2)-i-1; u++ {
			if p1[u] > p2[u+1] {
				x = p2[u]
				p2[u] = p2[u+1]
				p2[u+1] = x
			}
		}
	}
	fmt.Printf("sort p2 is %d\n", p2[:])
}
