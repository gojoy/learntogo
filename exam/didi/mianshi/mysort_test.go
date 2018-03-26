package mianshi

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func Test_qsort(t *testing.T) {
	list := make([]int, 0)
	for i := 0; i < 100; i++ {
		list = append(list, rand.Int())
		//sort.Ints(list)
		quickSort(list, 0, len(list)-1)
		if sort.IntsAreSorted(list) {
			t.Log("qosrt ok\n")
		} else {
			t.Error("qsort error\n")
			for _, v := range list {
				fmt.Printf("%v  ", v)
			}
			fmt.Println()
		}
	}
}
