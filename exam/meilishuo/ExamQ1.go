package meilishuo

import "fmt"

var longsmonths []int = []int{0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335}
var shortsmonths []int = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}

type t [3]int

func Q1() {
	var (
		y, m, d int
		err     error
	)
	s := make([]t, 0)
	_, err = fmt.Scanln(&y, &m, &d)
	for err == nil {
		s = append(s, t{y, m, d})
		//if isR(y) {
		//
		//	fmt.Println(longsmonths[m-1]+d)
		//} else {
		//	fmt.Println(shortsmonths[m-1]+d)
		//}
		_, err = fmt.Scanln(&y, &m, &d)
	}
	for _, v := range s {
		if isR(v[0]) {
			fmt.Println(longsmonths[m-1] + d)
		} else {
			fmt.Println(shortsmonths[m-1] + d)
		}
	}

}

func isR(y int) bool {

	if y%400 == 0 || y%100 != 0 && y%4 == 0 {
		//fmt.Println("run")
		return true
	}
	//fmt.Println("noyun")
	return false
}
