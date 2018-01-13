package shunfeng

import "fmt"

func Q1()  {
	var (
		k,a,b int
	)
	fmt.Scanln(&k,&a,&b)
	fmt.Println(q1(k,a,b))

}
func q1(k,a,b int) int  {
	var (
		count int
	)
	for i:=a;i<b;i++ {
		if i%k!=0 {
			continue
		}
		if f(i)*k==i {
			count++
		}
	}
	return count
}

func f(n int) int  {
	var (
		sum,yu int
	)
	for n!=0 {
		yu=n%10
		n=n/10
		sum+=yu*yu
	}
	return sum
}