package qunar

import "fmt"



func Q3()  {
	var (
		a,b,m uint64
	)
	fmt.Scanln(&a,&b,&m)
	fmt.Println(q3(a,b,m))
}
func q3(a,b,m uint64) uint64  {
	var (
		res uint64=1
		i uint64=0
	)
	for i=0;i<b;i++ {
		res=res*a
	}
	return res%m
}
