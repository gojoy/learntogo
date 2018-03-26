package pearls

//编程珠玑 随机数
import (
	"math/rand"
)

//返回0-n中m个随机数
func RandIntNum(m, n int) []int {
	j := 0
	res := make([]int, m)
	for i := 0; i < n; i++ {
		if m < 1 {
			break
		}
		if rand.Int()%(n-i) < m {
			//res=append(res,i)
			res[j] = i
			m--
			j++
			//fmt.Printf("i %v\n",i)
		}
	}
	return res
}

//使用map结构
func RandIntByMap(m, n int) []int {
	set := make(map[int]bool)
	res := make([]int, m)
	for i := 0; i < m; {
		t := rand.Int() % n
		if _, ok := set[t]; !ok {
			res[i] = t
			set[t] = true
			i++
		}
	}
	return res
}

//给定n面的骰子，每面概率在数组中，模拟摇骰子
func Roll(s []float32) int {
	l := len(s)
	temp := make([]float32, l-1)
	temp[0] = 0
	temp[l-2] = 1
	for i := 1; i < l-1; i++ {
		temp[i] = temp[i-1] + s[i]
	}
	//fmt.Printf("temp is %v\n",temp)
	r := rand.Float32()
	//fmt.Printf("r is %v\n",r)
	for i := 0; i < l-1; i++ {
		if r < temp[i] {
			return i
		}
	}
	return l - 1
}
