package wangyi

import "fmt"

func Q7()  {
	var (
		n,d,k,s int
	)
	fmt.Scanln(&n)
	sts:=make([]int,n)
	for i:=0;i<n;i++ {
		fmt.Scanf("%d",&s)
		sts[i]=s
	}
	nu,_:=fmt.Scanln(&k,&d)
	if nu!=2 {
		fmt.Scanln(&k,&d)
	}
	//fmt.Printf("num is %d,err is %v\n",nu,err)
	fmt.Printf("k is %d,s is %v\n",k,sts)
	fmt.Println(maxSubSubseqence(sts,d,k))
}

func maxSubSubseqence(s []int,d,k int) int  {
	if len(s)==1 {
		return s[0]
	}
	fmt.Printf("k is %d\n",k)
	l:=len(s)
	if l<k {
		//logg.Printf("k too big")
		return -1
	}

	//maxsub[y][x] 代表当一共有前x个同学，同时选y个同学时的最大乘积
	maxsub:=make([][]int,0)
	for i:=0;i<=k;i++ {
		n:=make([]int,l+1)
		maxsub=append(maxsub,n)
	}
	for i:=0;i<l+1;i++ {
		maxsub[0][i]=1
	}
	for i:=0;i<k+1;i++ {
		maxsub[i][0]=1
	}
	for i:=1;i<len(maxsub);i++ {
		maxsub[i][1]=s[0]
	}
	logg.Printf("x is %d,y is %d\n",len(maxsub),len(maxsub[0]))

	showMutliArray(maxsub)
	for i:=1;i<len(maxsub[0]);i++ {
		logg.Printf("i is %d\n",i)
		//if s[i-1]>maxsub[1][i-1] {
		//	maxsub[1][i]=s[i-1]
		//}
		//maxsub[1][i]=maxsub[1][i-1]
		maxsub[1][i]=s[i-1]
	}

	for kk:=2;kk<len(maxsub);kk++ {
		//logg.Printf("kk is %d\n",kk)
		for nn:=k;nn<len(maxsub[0]);nn++ {
			//logg.Printf("nn is %d\n",nn)
			selects:=make([]int,0)
			selects=append(selects,maxsub[kk-1][nn-1])
			for step:=1;step<d  ;step++ {
				if nn-step-1<0 {
					break
				}
				//logg.Printf("step is %d,y is %d,x is %d\n",step,kk-2,nn-step-1)
				selects=append(selects,maxsub[kk-2][nn-step-1]*s[nn-step-1]*s[nn-1])
			}
			maxsub[kk][nn]=maxslice(selects)
			//logg.Printf("maxsub x is %d,y is %d,value is %d\n",kk,nn,maxsub[kk][nn])
		}
	}

	//showMutliArray(maxsub)
	return maxslice(maxsub[len(maxsub)-1])
}

func showMutliArray(a [][]int)	  {
	y:=len(a)
	for i:=0;i<y;i++ {
		fmt.Println(a[i])
	}
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func maxslice(a []int) int {
	res:=a[0]
	for i:=1;i<len(a);i++ {
		if res<a[i] {
			res=a[i]
		}
	}
	return res
}