// t7.1array project main.go
package main

import (
	"fmt"
	"learntogo/t7.1array/list"
)

var (
	p1 [15]int //[15]int 型
	//new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：
	//这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体（参见第 10 章）；它相当于 &T{}。
	//未初始化
	p2     = new([15]int) //*[15]int型
	p      *int
	dp     [5][5]int
	slice1 []int
	slice2 []int
	//make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel
	//已经初始化
	//make(int,len,cap) len代表切片的长度，cap代表总长度，切片内加切片外
	slice3        = make([]int, 10, 15)
	str    string = "abcdefg"
	pi            = []int{1, 2, 3, 6, 40, 67, 7}
)

func testslice() {
	//*p2 = p1
	fmt.Printf("in main p1 type is %T,p2(new) type is %T,p type is %T.slice type is %T\n", p1, p2, p, slice1)
	fmt.Println("begin\n")
	//ptary(p2)
	//ptary(&p1)
	slice1 = p1[:]
	slice2 = p2[:]
	//ptary(p2)
	//getvalue(slice1)
	//ptary(p2)
	//outnopoint(slice1)
	//outnopoint(slice2)
	getvalue(slice1)
	//outnopoint(slice1)
	//slice3[] = slice1[0:9]
	slice3 = slice1[:]
	//fmt.Println(slice3[:])
	//outnopoint(slice3)
	slice3 = append(slice3, 17, 16, 19)
	fmt.Printf("type is %T,len is %v,cap is %v\n", slice3, len(slice3), cap(slice3))
	byteout("oknowtest", 4)
	//t713("abcdefg")
	t714fanz(str)
	fmt.Printf("type is %T,length is %d\n", pi, len(pi))
	mpsort(pi)
}

func main() {
	list.TestList()
}