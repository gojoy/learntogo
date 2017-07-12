//华为在线编程题目
package huawei


import "fmt"


//王强今天很开心，公司发给N元的年终奖。王强决定把年终奖用于购物，
// 他把想买的物品分为两类：主件与附件，附件是从属于某个主件的，
// 下表就是一些主件与附件的例子：
//主件	附件
//电脑	打印机，扫描仪
//书柜	图书
//书桌	台灯，文具
//工作椅	无
//如果要买归类为附件的物品，必须先买该附件所属的主件。每个主件可以有 0 个、 1 个或 2 个附件。
// 附件不再有从属于自己的附件。王强想买的东西很多，为了不超出预算，他把每件物品规定了一个重要度，
// 分为 5 等：用整数 1 ~ 5 表示，第 5 等最重要。他还从因特网上查到了每件物品的价格（都是 10 元的整数倍）。
// 他希望在不超过 N 元（可以等于 N 元）的前提下，使每件物品的价格与重要度的乘积的总和最大。
//设第 j 件物品的价格为 v[j] ，重要度为 w[j] ，共选中了 k 件物品，编号依次为 j 1 ， j 2 ，……， j k ，
// 则所求的总和为：
//v[j 1 ]*w[j 1 ]+v[j 2 ]*w[j 2 ]+ … +v[j k ]*w[j k ] 。（其中 * 为乘号）
//请你帮助王强设计一个满足要求的购物单。
func BuyThings()  {

}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

//切割钢条
//动态规划
func Divide() int {
	var p []int=[]int{0,1,5,8,9,10,17,17,20,24,30}
	n:=10
	fmt.Printf("n is %d,max is %d\n",n,cutBottomUp(n,p))


	return 0
}
//递归方法
func cut(n int,p []int) int {
	if n==0 {
		return 0
	}
	q:=0
	for i:=1;i<len(p)&&i<=n;i++ {
		q=max(q,p[i]+cut(n-i,p))
		fmt.Printf("i is %d,q is %d\n",i,q)
	}
	//fmt.Printf("p is %d\n",q)
	return q
}

//非递归方法
func cutBottomUp(n int,p []int) int {
	var (
		l []int
	)
	r:=make([]int,n+1)
	res:=make([][]int,n+1)
	r[0]=0

	for j:=1;j<=n;j++ {
		q:=0
		tmpl:=0
		var i int
		for i=1;i<=j && i<len(p);i++ {
			if q<p[i]+r[j-i] {
				q=p[i]+r[j-i]
				tmpl=i
			}
		}
		l=append(res[j-tmpl],tmpl)
		//fmt.Printf("j is %d,tmpl is %d,last is %v,l is %v\n",j,tmpl,res[j-tmpl],l)
		res[j]=make([]int,len(l))
		copy(res[j],l)
		//fmt.Printf("res is %v\n",res)
		r[j]=q
	}
	fmt.Printf("r is %v\n,res is %v\n",r,res[1:])
	return r[n]
}

//最长公共子序列
//Z=<B,C,D,B> X=<A,B,C,B,D,A> 则Z为X的子序列 只需要出现顺序相同，不需要相邻
func LCS()   {
	var (
		X string="ABCBDAB"
		Y string="BDCABA"
		//X string="A"
		//Y string="A"
	)
	//fmt.Printf("lcs is %d\n",getlcsrec(X,Y,len(X),len(Y)))
	fmt.Printf("res is %v\n",getLCSByAry(X,Y))
}

func getlcsrec(x,y string, i,j int) int {
	if i<0 || j <0 {
		panic("error < 0")
	}
	if i==0 || j==0 {
		return 0
	}
	if  x[i-1]==y[j-1] {
		return getlcsrec(x,y,i-1,j-1)+1
	}
	return max(getlcsrec(x,y,i-1,j),getlcsrec(x,y,i,j-1))
}

func getLCSByAry(x,y string) int {
	lx:=len(x)
	ly:=len(y)
	if lx==0 || ly==0 {
		return 0
	}

	res:=make([][]int,lx+1)
	for i:=0;i<=lx;i++ {
		res[i]=make([]int,ly+1)
	}
	for i:=1;i<=lx;i++ {
		for j:=1;j<=ly;j++ {
			if x[i-1]==y[j-1] {
				res[i][j]=res[i-1][j-1]+1
			} else {
				res[i][j] = max(res[i - 1][j], res[i][j - 1])
			}
		}
	}
	fmt.Printf("r is %v\n",res)
	return res[lx][ly]
}