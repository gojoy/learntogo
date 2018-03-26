package qunar

import (
	"fmt"
	"strconv"
	"strings"
)

type node [3]int

func Q2() {
	var (
		root, val, left, right, n int
		err                       error
		str                       string
		flag                      int = 1
	)
	trees := make([]node, 0)
	fmt.Scanln(&root)
	n, err = fmt.Scanln(&str)
	for err == nil && n == 1 {
		//fmt.Printf("input str is %v\n",str)
		if str[len(str)-1] == '\n' {
			str = str[:len(str)-1]
		}
		strs := strings.Split(str, ":")
		//fmt.Printf("strs is %v\n",strs)
		val, _ = strconv.Atoi(strs[0])
		leaf := strings.Split(strs[1], "|")
		left, _ = strconv.Atoi(leaf[0])
		right, _ = strconv.Atoi(leaf[1])
		trees = append(trees, node{val, left, right})
		n, err = fmt.Scanln(&str)
	}
	//fmt.Printf("trees is %v\n",trees)
	res := make([]int, 0)
	rootpos := findnode(root, trees)
	midtree(rootpos, trees, &res)
	//fmt.Printf("res is %v\n",res)
	for i := 0; i < len(res)-2; i++ {
		if res[i] > res[i+1] {
			flag = 0
			break
		}
	}
	fmt.Println(flag)
}

func findnode(l int, trees []node) int {
	for i := 0; i < len(trees); i++ {
		if trees[i][0] == l {
			return i
		}
	}
	return -1
}

func midtree(pos int, trees []node, list *[]int) {
	leftnode := trees[pos][1]
	if leftnode != -1 {
		midtree(findnode(leftnode, trees), trees, list)
	}
	*list = append(*list, trees[pos][0])
	rightnode := trees[pos][2]
	if rightnode != -1 {
		midtree(findnode(rightnode, trees), trees, list)
	}
}
