package pearls

//编程珠玑 随机数
import (
	"math/rand"
)
//返回0-n中m个随机数
func RandIntNum(m, n int) []int	 {
	j:=0
	res:=make([]int,m)
	for i:=0;i<n;i++ {
		if m<1 {
			break
		}
		if(rand.Int()%(n-i)<m) {
			//res=append(res,i)
			res[j]=i
			m--
			j++
			//fmt.Printf("i %v\n",i)
		}
	}
	return res
}

//使用map结构
func RandIntByMap(m,n int) []int  {
	set:=make(map[int]bool)
	res:=make([]int,m)
	for i:=0;i<m; {
		t:=rand.Int()%n
		if _,ok:=set[t];!ok {
			res[i]=t
			set[t]=true
			i++
		}
	}
	return res
}