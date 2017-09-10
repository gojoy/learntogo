package wangyi

import "fmt"

func EXAMQ3()  {
	var (
		n,l,c int
	)
	fmt.Scanln(&n,&l)
	p:=make([]int,n-1)
	isarrive:=make([]bool,n)
	for i:=0;i<len(isarrive);i++ {
		isarrive[i]=false
	}
	isarrive[0]=true
	for i:=0;i<n-1;i++ {
		fmt.Scan(&c)
		p[i]=c
	}
	fmt.Println(youli(l,0,1,p,isarrive))
}


func youli(l int,nowcity int,lookcities int,citylist []int,isarrive []bool) int  {
	if l==0 {
		return lookcities
	}
	max:=lookcities
	res:=lookcities
	for i:=0;i<len(citylist);i++ {
		//第i+1城市与当前城市相连，选择它
		if citylist[i]==nowcity {
			//已结去过该城市
			if isarrive[i+1] {
				res=youli(l-1,i+1,lookcities,citylist,isarrive)
			} else {
				isarrive[i+1]=true
				res=youli(l-1,i+1,lookcities+1,citylist,isarrive)
			}
			if res>max {
				max=res
			}
		}
	}
	return max
}

func maxyouli(isarrive []bool,citylist []int,nowcity int,l int) int  {
	max:=0
	for i:=0;i<len(citylist);i++ {
		if citylist[i]==nowcity {
			isarrive[i+1]=true
			res:=youli(l-1,i+1,2,citylist,isarrive)
			if res>max {
				max=res
			}
		}
	}
	return max
}