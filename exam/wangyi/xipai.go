package wangyi

import (
	"fmt"
	"log"
	"os"
)

var logg = log.New(os.Stdout, "log:", log.Lshortfile)


func Input() {
	var (
		T,m,n,k,c int
		err error
	)
	m,err=fmt.Scanln(&T)
	if err!=nil {
		logg.Printf("m is %v,err is %v",m,err)
	}
	for i:=0;i<T;i++ {
		m,err=fmt.Scanln(&n,&k)
		if err!=nil {
			logg.Printf("m is %v,err is %v",m,err)
		}
		cards:=make([]int,n*2)
		for j:=0;j<n*2;j++ {
			m,err=fmt.Scanf("%d",&c)
			if err!=nil {
				logg.Printf("m is %v,err is %v",m,err)
			}
			cards[j]=c
		}
		//fmt.Printf("T is %d, cards is %v\n",T,cards)
		xipai(cards,k)

	}
	return
}

//洗牌问题
func xipai(cards []int, k int) {
	var (
		length int=len(cards)
		//n int=length/2
	)
	tmpcards:=make([]int,length)

	for i:=0;i<k;i++ {
		for j:=1;j<=length;j++ {
			reali:=length-j+1
			tmppos:=(2*reali)%(length+1)-1
			//logg.Printf("reali is %v,j is %d,tmppos is %d\n",reali,j,tmppos)
			tmpcards[tmppos]=cards[j-1]
		}
		copy(cards,tmpcards)
		swapnext(cards)
		swapcards(cards)
		//fmt.Printf("cards is %v\n",cards)
	}
	output(cards)
}

func swapnext(a []int) {
	l:=len(a)
	if l<2 || l%2!=0 {
		logg.Printf("swapnext error\n")
		return
	}
	for i:=1;i<l; {
		a[i],a[i-1]=a[i-1],a[i]
		i=i+2
	}
}

func swapcards(a []int){
	l:=len(a)
	for i:=0;i<l/2;i++ {
		a[i],a[l-i-1]=a[l-i-1],a[i]
	}
	return
}

func output(a []int)  {
	var i int
	for i=0;i<len(a)-1;i++ {
		fmt.Printf("%d ",a[i])
	}
	fmt.Printf("%d\n",a[i])
}

