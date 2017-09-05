package wangyi

import (
	"fmt"
	"strings"
	"math"
)

func Tangg() {
	var (
		n1,n2,n3,n4,n int
		err error
		flag bool=true
	)
	n,err=fmt.Scanln(&n1,&n2,&n3,&n4)
	if err!=nil || n!=4 {
		return
	}

	if (n1+n3)%2!=0 || (n2+n4)%2!=0 {
		flag=false
	}
	a:=(n1+n3)/2
	b:=(n2+n4)/2
	c:=n4-b
	if a-b!=n1 || b-c!=n2 || a+b!=n3 || !flag {
		fmt.Printf("No")
	} else {
		fmt.Printf("%d %d %d",a,b,c)
	}
}


//一个只包含'A'、'B'和'C'的字符串，如果存在某一段长度为3的连续子串中恰好'A'、'B'和'C'各有一个，那么这个字符串就是纯净的，否则这个字符串就是暗黑的。例如：
//BAACAACCBAAA 连续子串"CBA"中包含了'A','B','C'各一个，所以是纯净的字符串
//AABBCCAABB 不存在一个长度为3的连续子串包含'A','B','C',所以是暗黑的字符串
//你的任务就是计算出长度为n的字符串(只包含'A'、'B'和'C')，有多少个是暗黑的字符串。
func Anhei()  {
	var (
		n int
		err error
	)
	_,err=fmt.Scanln(&n)
	if err!=nil {
		return
	}
	a:=make([]int,n+1)
	a[0]=0
	a[1]=3
	a[2]=9
	for i:=3;i<n+1;i++ {
		//递推公式 f(n)=f(n-1)*2+f(n-2)
		//通过生成树可以确定公式
		a[i]=2*a[i-1]+a[i-2]
	}
	fmt.Printf("%d",a[n])
}

//牛牛需要每次都回答 t 是否是 s 的子序列。
// 注意，子序列不要求在原字符串中是连续的，
// 例如串 abc，它的子序列就有 {空串, a, b, c, ab, ac, bc, abc} 8 种。
func Cangbaotu()  {
	var (
		s,t string
	)
	fmt.Scanln(&s)
	fmt.Scanln(&t)
	if len(t)<1 {
		fmt.Printf("Yes")
		return
	}
	if zixulie(s,t) {
		fmt.Printf("Yes")
		return
	}
	fmt.Printf("No")

}

//t是否是s的子序列
func zixulie(s,t string) bool  {
	pos:=0
	tmp:=0
	for i:=0;i<len(t);i++ {
		//fmt.Printf("i is %d,ti is %c,pos is %v,tmp is %v,spos is %v\n",i,t[i],pos,tmp,s[tmp:])
		pos=strings.IndexByte(s[tmp:],t[i])
		tmp=tmp+pos+1
		if pos==-1 || pos>=len(s) {
			return false
		}
	}
	return true
}

//航天飞行器是一项复杂而又精密的仪器，飞行器的损耗主要集中在发射和降落的过程，
// 科学家根据实验数据估计，如果在发射过程中，产生了 x 程度的损耗，那么在降落的过程中就会产生 x2 程度的损耗，
// 如果飞船的总损耗超过了它的耐久度，飞行器就会爆炸坠毁。问一艘耐久度为 h 的飞行器，假设在飞行过程中不产生损耗，
// 那么为了保证其可以安全的到达目的地，只考虑整数解，至多发射过程中可以承受多少程度的损耗？

func Fashe()  {
	var h int64
	fmt.Scanln(&h)
	fmt.Println(fashe(h))
}

func fashe(h int64) int64{

	res:=math.Sqrt(float64(h)+float64(0.25))-0.5
	return int64(res)
}

//给你一个N，你想让其变为一个Fibonacci数，每一步你可以把当前数字X变为X-1或者X+1，
// 现在给你一个数N求最少需要多少步可以变为Fibonacci数。
func TranFibonacci()  {
	var (
		n int
		i int
	)
	fmt.Scanln(&n)
	t:=make([]int,100)
	t[0]=0
	t[1]=1
	for i=2;i<100;i++ {
		t[i]=t[i-1]+t[i-2]
		if t[i]>=n {
			break
		}
	}
	left:=n-t[i-1]
	right:=t[i]-n
	if left<right {
		fmt.Println(left)
		return
	}
	fmt.Println(right)
}


//有一片1000*1000的草地，小易初始站在(1,1)(最左上角的位置)。
// 小易在每一秒会横向或者纵向移动到相邻的草地上吃草(小易不会走出边界)。
// 大反派超超想去捕捉可爱的小易，他手里有n个陷阱。第i个陷阱被安置在横坐标为xi ，
// 纵坐标为yi 的位置上，小易一旦走入一个陷阱，将会被超超捕捉。你为了去解救小易，
// 需要知道小易最少多少秒可能会走入一个陷阱，从而提前解救小易
func Savexiaoyi()  {
	var (
		n,x,y int
	)
	fmt.Scanln(&n)
	xs:=make([]int,n)
	ys:=make([]int,n)
	for i:=0;i<n;i++ {
		fmt.Scan(&x)
		xs[i]=x
	}
	for i:=0;i<n;i++ {
		fmt.Scan(&y)
		ys[i]=y
	}
	fmt.Println(savexiaoyi(xs,ys))
}

func savexiaoyi(xs,ys []int) int  {
	l:=len(xs)
	min:=10000
	for i:=0;i<l;i++ {
		cost:=xs[i]+ys[i]-2
		if cost<min {
			min=cost
		}
	}
	return min
}