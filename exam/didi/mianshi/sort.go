package mianshi

import "fmt"

func insertSort(a []int){
	var (
		i,j,key int
	)
	l:=len(a)
	if len(a)<2 {
		return
	}
	for i=1;i<l;i++ {
		key=a[i]
		for j=i-1;j>=0 && a[j]>key;j-- {
			a[j+1]=a[j]
		}
		a[j+1]=key
	}
}

func bubbleSort(a []int)  {
	l:=len(a)
	var (
		i,j int
	)
	for i=0;i<l;i++ {
		for j=1;j<l-i;j++ {
			if a[j]<a[j-1] {
				a[j],a[j-1]=a[j-1],a[j]
			}
		}
	}
	return 
}

func quickSort(a []int,left,right int)  {
	var (
		p,i,j int
	)
	if left>right {
		return
	}
	p=a[left]
	i=left
	j=right
	for i!=j {
		for a[j]>=p && i<j {
			j--
		}
		for a[i]<=p && i<j {
			i++
		}
		if i<j {
			a[i],a[j]=a[j],a[i]
		}
	}
	a[left]=a[i]
	a[i]=p
	quickSort(a,left,i-1)
	quickSort(a,i+1,right)
}

func Tqsort()  {

	l:=[]int{4,2,7,8,1,5,3,6,9,0}
	fmt.Printf("l is %v\n",l)
	quickSort(l,0,len(l)-1)

	fmt.Printf("after sort l is %v\n",l)
}
