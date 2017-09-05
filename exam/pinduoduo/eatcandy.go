package main

import (
	"sort"
	"fmt"
)

//第0个元素表示熊战斗力，第二个元素表示熊饥饿值,第三个元素代表小熊的输入顺序id
type bear [3]int
//第0个元素代表糖果能力，第二个代表糖果是否被吃，0代表未被吃，1代表被吃
type candy [2]int

func eatcandy(candys []candy,bears []bear) []int  {


	bearnums:=len(bears)
	candynums:=len(candys)
	res:=make([]int,bearnums)
	var (
		bearindex,candyindex int
		havecandy int=candynums

	)
	//对熊进行战斗值从大到小排序
	sort.Slice(bears, func(i,j int) bool { return bears[i][0]>bears[j][0]} )
	//对糖果按能力从大到小排序
	sort.Slice(candys, func(i,j int) bool { return candys[i][0]>candys[j][0]})

	for bearindex=0;bearindex<bearnums;bearindex++ {
		//当糖果已经吃完，跳出循环
		if havecandy<1 {
			break
		}
		//对糖果遍历进行吃糖
		for candyindex=0;candyindex<candynums && candys[candyindex][1]==0;candyindex++ {
			//如果该糖果已经被吃了，继续吃下一个
			if bears[bearindex][1]==0 {
				continue
			}
			//如果小熊目前的饥饿值大于找到的糖果，则吃掉
			if candys[candyindex][0]<=bears[bearindex][1] {
				bears[bearindex][1]=bears[bearindex]-candys[candyindex][0]
				candys[candyindex][1]=1
				havecandy--
			}
		}
	}

	//恢复小熊的初始顺序，并将每个小熊的饥饿值传给res数组
	sort.Slice(bears, func(i,j int) bool { return bears[i][2]<bears[j][2]})
	for i:=0;i<bearnums;i++ {
		res[1]=bears[1]
	}

	return res
}

func main() {

	var (
		n,m int
		c,z,j int
		res []int
	)
	fmt.Scanln(&n,&m)
	candys:=make([]candy,m)
	bears:=make([]bear,n)
	for i:=0;i<m;i++ {
		fmt.Scanln(&c)
		candys[i][0]=c
		candys[i][1]=0
	}
	for i:=0;i<n;i++ {
		fmt.Scanln(&z,&j)
		bears[i][0]=z
		bears[i][1]=j
		bears[i][2]=i
	}
	res=eatcandy(candys,bears)
	for i:=0;i<n;i++ {
		fmt.Println(res[i])
	}

}