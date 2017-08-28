package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

var logg = log.New(os.Stdout, "log:", log.Lshortfile)

//PMidea 0:id 1:创建时间 2:优先级 3:需要时间 4:是否已交付程序员处理 5：预计完成时间 6:idea id
type idea [7]int

type programmer struct {
	busy        bool
	workEndtime int
}

func main() {
	var (
		//N个PM，M个程序员，P个idea
		N, M, P                                         int
		pmId, ideaCreatTime, ideaPriority, ideaNeedTime int
	)
	n, err := fmt.Scanf("%d %d %d", &N, &M, &P)
	if err != nil || n != 3 {
		fmt.Printf("scanf num is %d error: %v\n", n, err)
		return
	}

	pms := make([]programmer, M)
	for i := 0; i < M; i++ {
		pms[i].busy = false
		pms[i].workEndtime = 0
	}

	ideas := make([]idea, P)
	for i := 0; i < P; i++ {
		n, err := fmt.Scanf("%d %d %d %d", &pmId, &ideaCreatTime, &ideaPriority, &ideaNeedTime)
		if err != nil || n != 4 {
			fmt.Printf("scanf num is %d idea error:%v\n", n, err)
		}
		ideas[i] = idea{pmId, ideaCreatTime, ideaPriority, ideaNeedTime, 0, 0, i}
		//ideas=append(ideas,idea{pmId,ideaCreatTime,ideaPriority,ideaNeedTime})
	}
	fmt.Printf("idea is %v\n", ideas)
	doWorkTime(pms, ideas, N)
	return
}

func doWorkTime(pgs []programmer, ideas []idea, N int) int {

	if len(ideas) == 0 || len(pgs) == 0 {
		return 0
	}
	var (
		time  int = -1
		still bool=false
	)

	for {
		time++
		doWork(pgs, time)
		if still == true {
			logg.Printf("all idea have done ideas is %v\n",ideas)
			break
		}
		for j := 0; j < N; j++ {
			still = true
			for i := 0; i < len(ideas); i++ {
				if ideas[i][4] == 0 {
					still = false && still
				} else {
					still = true && still
				}
			}

			nextidea, ok := selectIdea(ideas, time)
			if !ok {
				fmt.Printf("now %d have no idea to do\n", time)
				continue
			}
			fmt.Printf("next idea is %v\n", ideas[nextidea])
			if !handoutIdea(ideas, nextidea, pgs, time) {
				fmt.Printf("now %d have no pg to do\n", time)
				continue
			}
			logg.Printf("have done idea %v\n", ideas[nextidea])

		}
	}
	for i := 0; i < len(ideas); i++ {
		fmt.Printf("%d idea finish time is %d\n", ideas[i][0], ideas[i][5])
	}
	return time

}

func handoutIdea(ideas []idea, id int, pgs []programmer, time int) bool {

	for i := 0; i < len(pgs); i++ {
		if !pgs[i].busy {
			pgs[i].workEndtime = time + ideas[id][3]
			pgs[i].busy = true
			ideas[id][4] = 1
			ideas[id][5] = time + ideas[id][3]
			logg.Printf("handout idea is %v\n", ideas[id])
			return true
		}
	}
	return false
}

func selectIdea(ideas []idea, time int) (int, bool) {
	logg.Printf("selectidea time is %v\n ideas is %v\n", time, ideas)
	ids := make([]idea, 0)
	samepriority := make([]idea, 0)
	for i := 0; i < len(ideas); i++ {
		if ideas[i][1] <= time && ideas[i][4] == 0 {
			ids = append(ids, ideas[i])
		}
	}
	logg.Printf("ids is %v\n", ids)

	if len(ids) == 0 {
		return ideas[0][6], false
	}
	if len(ids) == 1 {
		return ids[0][6], true
	}
	sort.SliceStable(ids, func(i, j int) bool { return ids[i][2] > ids[j][2] })
	samepriority = append(samepriority, ids[0])
	for i := 1; i < len(ids); i++ {
		if ids[i][2] != ids[0][2] {
			break
		}
		samepriority = append(samepriority, ids[i])
	}
	logg.Printf("same is %v\n", samepriority)
	if len(samepriority) == 1 {
		return samepriority[0][6], true
	}
	sort.Slice(samepriority, func(i, j int) bool { return samepriority[i][0] < samepriority[j][0] })
	return samepriority[0][6], true
}

func doWork(pgs []programmer, time int) {
	for i := 0; i < len(pgs); i++ {
		if pgs[i].workEndtime == time {
			pgs[i].busy = false
			pgs[i].workEndtime = 0
		}
	}
	return
}
