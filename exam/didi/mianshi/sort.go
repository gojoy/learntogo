package mianshi

func insertSort(a []int)  {
	l:=len(a)
	if len(a)<2 {
		return
	}
	for i:=2;i<l-1;i++ {
		key:=a[i]
		j:=i-1
		for j>0 && a[j]>key {
			a[j+1]=a[j]
			j--
		}
		a[j+1]=key
	}
}

func bubbleSort(a []int)  {
	l:=len(a)
	for i:=0;i<l-1;i++ {
		for j:=i+1;j<l;j++ {
			if a[j]<a[j-1] {
				a[j],a[j-1]=a[j-1],a[j]
			}
		}
	}
	return 
}

func quickSort(a []int,left,right int)  {
	if left==right {
		return
	}
	p:=a[left]
	i:=left
	j:=right
	for i<j {
		for a[j]<p && i<j {
			j--
		}
		for a[i]>p && i<j {
			i++
		}
		if i<j {
			a[i],a[j]=a[i],a[j]
		}
	}
	a[left]=a[i]
	a[i]=p
	quickSort(a,left,i-1)
	quickSort(a,i+1,right)

}
