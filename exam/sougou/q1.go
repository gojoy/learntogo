package sougou

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Q1() {
	var (
		n   int
		num float64
	)
	fmt.Scanln(&n)
	points := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&num)
		points[i] = num

		fmt.Printf("%.8f", maxdistance1(points))

	}
}

func Q1buf() {
	var (
		n   int
		num float64
		err error
	)
	fmt.Scanln(&n)
	points := make([]float64, n)
	inputReader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		s, _ := inputReader.ReadString('\n')
		//fmt.Printf("s is %v\n",s)
		num, err = strconv.ParseFloat(s[:len(s)-1], 64)
		if err != nil {
			fmt.Println(err)
		}
		points[i] = num
	}
	fmt.Printf("%.8f", maxdistance1(points))
}

func maxdistance(points []float64) float64 {
	var (
		dis, rui, start float64
		max             float64
		//flag bool=false
	)
	start = points[0]
	max = 0
	if start <= 180 {
		rui = start + 180
		pos := sort.SearchFloat64s(points, rui)
		if points[pos] > rui {
			dis = 360 - points[pos] + start
			max = points[pos-1] - start
			if dis > max {
				max = dis
			}
		} else {
			max = points[pos] - start
			dis = 360 - points[pos+1] + start
			if dis > max {
				max = dis
			}
		}

		//for i:=1;i<len(points);i++ {
		//	if points[i]<=rui {
		//		dis=points[i]-start
		//	} else {
		//		if flag {
		//			break
		//		}
		//		dis=360-points[i]+start
		//		flag=true
		//	}
		//	if dis>max {
		//		max=dis
		//	}
		//}
	} else {
		//for i:=1;i<len(points);i++ {
		//	dis=points[i]-start
		//	if dis>max {
		//		max=dis
		//	}
		//}
		max = points[len(points)-1] - start
	}
	return max
}

func maxdistance1(points []float64) float64 {
	var (
		dis, rui, start float64
		max             float64
		flag            bool = false
	)
	start = points[0]
	max = 0
	if start <= 180 {
		rui = start + 180
		for i := 1; i < len(points); i++ {
			if points[i] <= rui {
				dis = points[i] - start
			} else {
				if flag {
					break
				}
				dis = 360 - points[i] + start
				flag = true
			}
			if dis > max {
				max = dis
			}
		}
	} else {
		//for i:=1;i<len(points);i++ {
		//	dis=points[i]-start
		//	if dis>max {
		//		max=dis
		//	}
		//}
		max = points[len(points)-1] - start
	}
	return max
}
