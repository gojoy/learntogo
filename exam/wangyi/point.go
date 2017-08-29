package wangyi

import (
	"fmt"
	"math"
)

//小易有一个圆心在坐标原点的圆，小易知道圆的半径的平方。小易认为在圆上的点而且横纵坐标都是整数的点是优雅的，小易现在想寻找一个算法计算出优雅的点的个数，请你来帮帮他。
//例如：半径的平方如果为25
//优雅的点就有：(+/-3, +/-4), (+/-4, +/-3), (0, +/-5) (+/-5, 0)，一共12个点。
func Q2()  {
	var (
		r int
	)
	fmt.Scanln(&r)
	fmt.Println(getPoint(r))

}

func getPoint(rr int) int  {
	var (
		rint int
		x int
		y float64
		yy int
		res int
	)
	frr:=float64(rr)
	r:=math.Sqrt(frr)
	rint=int(r)
	if !isint(r) {
		rint++
	}
	for x=0;x<rint;x++ {
		yy=rr-x*x
		y=math.Sqrt(float64(yy))
		//fmt.Printf("x is %d,y is %v\n",x,y)
		if isint(y) {
			res++
			fmt.Printf("x is %d,y is %v\n",x,y)
		}
	}
	return res*4
}

func isint(y float64)bool  {
	it:=int(y)
	ry:=float64(it)
	//fmt.Printf("ty is %v\n",ry)
	if y-ry==0 {
		return true
	}
	return false
}
