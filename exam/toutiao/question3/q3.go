package main

import "fmt"

//idea 0:id 1:创建时间 2:优先级 3:需要时间 4:是否已交付程序员处理 5：预计完成时间
type idea [5]int

type programmer struct {
	busy bool
	workEndtime int
}

func main() {
	var (
		//N个PM，M个程序员，P个idea
		N,M,P int
		pmId,ideaCreatTime,ideaPriority,ideaNeedTime int
	)
	n,err:=fmt.Scanf("%d %d %d",&N,&M,&P)
	if err!=nil || n!=3 {
		fmt.Printf("scanf num is %d error: %v\n",n,err )
		return
	}

	pms:=make([]programmer,M)
	for i:=0;i<M;i++ {
		pms[i].busy=false
		pms[i].workEndtime=0
	}

	ideas:=make([]idea,P)
	for i:=0;i<P;i++ {
		n,err:=fmt.Scanf("%d %d %d %d",&pmId,&ideaCreatTime,&ideaPriority,&ideaNeedTime)
		if err!=nil || n!=4 {
			fmt.Printf("scanf num is %d idea error:%v\n",n,err)
		}
		ideas[i]=idea{pmId,ideaCreatTime,ideaPriority,ideaNeedTime,0}
		//ideas=append(ideas,idea{pmId,ideaCreatTime,ideaPriority,ideaNeedTime})
	}
	fmt.Printf("idea is %v\n",ideas)
	return
}

func doWorkTime(pgs []programmer,ideas []idea) int {

	if len(ideas)==0 || len(pgs)==0{
		return 0
	}
	var (
		time int
		still bool
	)

	for {
		still=true
		for i:=0;i<len(i);i++ {
			if ideas[i][4]==0 {
				still=false && still
				
			} else {
				still=true && still
			}
		}
		time++
		if still==true {
			break
		}
	}
	return time
}

func handoutIdea(ide idea,pgs []programmer,time int) bool  {

	for i:=0;i<len(pgs);i++ {
		if !pgs[i].busy {
			pgs[i].workEndtime=time+ide[3]
			pgs[i].busy=true
			return true
		}
	}
	return false
}

func selectIdea(ideas []idea,time int) idea {
	ids:=make([]idea,0)
	for i:=0;i<len(ideas);i++ {
		if ideas[i][1]>=time {
			ids=append(ids,ideas[i])
		}
	}
}