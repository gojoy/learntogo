package wangyi

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//牛牛想尝试一些新的料理，每个料理需要一些不同的材料，问完成所有的料理需要准备多少种不同的材料。
func Q8() {
	var (
		foods  map[string]bool = make(map[string]bool)
		counts int
	)

	input := bufio.NewReader(os.Stdin)
	s, err := input.ReadString('\n')
	for err == nil && s[0] != 10 {
		//logg.Printf("s is %d\n",s[0])
		strs := strings.Split(s[:len(s)-1], " ")
		//fmt.Printf("strs is %v\n",strs)
		counts = counts + count(foods, strs)
		s, err = input.ReadString('\n')
	}
	fmt.Printf("%d\n", counts)
	//fmt.Println(err)
}

func Q8byscanf() {
	var (
		s      string
		counts int
		err    error
	)
	foods := make(map[string]bool)
	_, err = fmt.Scan(&s)
	for err == nil {
		//fmt.Printf("s is %v\n",s)
		counts = counts + countsstring(foods, s)
		_, err = fmt.Scan(&s)
	}
	fmt.Println(counts)
	//fmt.Printf("n is %v,err is %v\n",n,err)
}

func countsstring(f map[string]bool, s string) int {
	var (
		res int
	)
	if _, ok := f[s]; !ok {
		res++
		f[s] = true
	}
	return res
}

func count(f map[string]bool, s []string) int {
	var (
		res int
		ok  bool
	)
	for i := 0; i < len(s); i++ {
		if s[i] == " " {
			continue
		}
		if _, ok = f[s[i]]; ok {

		} else {
			f[s[i]] = true
			res++
		}
	}
	return res
}
