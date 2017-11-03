package yy

import (
	"container/list"
	"fmt"
	"sort"
)

type component [3]int

func Q1() {
	var (
		start, end, v int
		err           error
	)
	coms := make([]component, 0)
	_, err = fmt.Scanln(&start, &end, &v)
	for err == nil {
		coms = append(coms, component{start, end, v})
		_, err = fmt.Scanln(&start, &end, &v)
	}
	fmt.Printf("err is %v,coms is %v\n\n", err, coms)
	deal(coms)
}

func deal(all []component) {
	var (
		s, i int
	)
	sort.Slice(all, func(i, j int) bool { return all[i][0] < all[j][0] })
	sm := make(map[int]int)
	em := make(map[int]int)
	for i := 0; i < len(all); i++ {
		sm[all[i][0]] = all[i][2]
		em[all[i][1]] = all[i][2]
	}

	ids := list.New()
	ids.PushBack(all[0][2])
	s = all[0][0]
	for i = all[0][0]+1; i <= all[len(all)-1][1]; i++ {
		//fmt.Printf("i is %v\n",i)
		if id, ok := sm[i]; ok {

			fmt.Printf("%d %d", s, i)
			printlist(ids)
			ids.PushBack(id)
			s = i
		}
		if id, ok := em[i]; ok {

			fmt.Printf("%d %d", s, i)
			printlist(ids)
			ids.Remove(getelem(id, ids))
			s = i
		}
	}
}

func printlist(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf(" %d", e.Value)
	}
	fmt.Printf("\n")
}

func getelem(v int, l *list.List) *list.Element {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == v {
			return e
		}
	}
	fmt.Printf("cannot find %v\n",v)
	return nil
}
