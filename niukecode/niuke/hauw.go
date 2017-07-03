package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"strconv"
	"log"
	"os"

	"bufio"

)

var logg=log.New(os.Stderr,"log:", log.Lshortfile)


//写出一个程序，接受一个有字母和数字以及空格组成的字符串，
// 和一个字符，然后输出输入字符串中含有该字符的个数。
// 不区分大小写。
func StringCountWord(s string, c uint8) (int, error) {
	var count int = 0
	if len(s) == 0 {
		return 0, errors.New("nil string")
	}
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			count++
		}
	}
	return count, nil
}

//明明想在学校中请一些同学一起做一项问卷调查，为了实验的客观性，
// 他先用计算机生成了N个1到1000之间的随机整数（N≤1000），
// 对于其中重复的数字，只保留一个，把其余相同的数去掉，不同的数对应着不同的学生的学号。
// 然后再把这些数从小到大排序，按照排好的顺序去找同学做调查。
// 请你协助明明完成“去重”与“排序”的工作
func RandNum() []int {
	var num int
	fmt.Scanf("%d", &num)

	//fmt.Printf("num is %d\n",num)
	input := make([]int, num)
	for i := 0; i < len(input); i++ {
		//fmt.Printf("input %d:\n",i)
		fmt.Scanf("%d", &input[i])
	}
	//fmt.Printf("input is %v\n",input)
	sort.Ints(input)
	//fmt.Printf("after sort is %v\n",input)
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			input = append(input[:i], input[i+1:]...)
			i++
		}
	}
	//fmt.Printf("now is %v\n",input)
	return input
}

//•连续输入字符串，请按长度为8拆分每个字符串后输出到新的字符串数组；
//•长度不是8整数倍的字符串请在后面补数字0，空字符串不处理
func StringSpli() {
	var (
		str1 string
		str2 string
	)
	fmt.Scanf("%s",&str1)
	fmt.Scanf("%s",&str2)

	strsp := func(str1 string) {
		if len(str1) == 0 {
			return
		}
		var i int
		length := ((len(str1)-1)/8 + 1) * 8
		//fmt.Println(length)
		res1 := make([]uint8, length)
		for i = 0; i < length; i++ {
			if i < len(str1) {
				res1[i] = str1[i]
			} else {
				res1[i] = '0'
			}
		}

		for j := 0; j < length; j++ {
			if j%8 == 0 && j > 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("%c", res1[j])
		}
	}

	strsp(str1)
	fmt.Printf("\n")
	strsp(str2)

}
//•连续输入字符串，请按长度为8拆分每个字符串后输出到新的字符串数组；
//•长度不是8整数倍的字符串请在后面补数字0，空字符串不处理
func Split() {

	f:=func(str string) {
		for len(str) >= 8 {
			fmt.Println(str[:8])
			str = str[8:]
		}
		if len(str) != 0 && len(str) < 8 {
			str = str + "00000000"
			fmt.Println(str[:8])
		}
	}
	var s string
	var s1 string
	fmt.Scanf("%s",&s)
	fmt.Scanf("%s",&s1)
	f(s)
	f(s1)
}

func TransSystem() {

	var str1 string
	fmt.Scanf("%s",&str1)
	num,err:=strconv.ParseInt(str1,0,0)
	if err != nil {
		fmt.Println(err)
		return
	}
	res:=strconv.Itoa(int(num))
	fmt.Println(res)
}

//将输入十六进制输出十进制
func TranSys1() {
	var (
		a int
		n int
		err error
	)
	n,err=fmt.Scanf("%x",&a)
	//fmt.Printf("n is %d,err is %v\n",n,err)
	for err == nil && n>0 {
		fmt.Printf("%d\n",a)
		n,err=fmt.Scanf("%v",&a)
	}
	fmt.Println(err)

}

//功能:输入一个正整数，按照从小到大的顺序输出它的所有质数的因子（如180的质数因子为2 2 3 3 5 ）
func getResult(uiDataInput int64)   {
	nextnum:=func(pre int64) int64 {
		if pre==2 {
			return 3
		}
		var j int64=3
		pre+=2
		for j < pre / 2+1 {
			if pre%j==0 {
				pre+=2
				j=2
			} else {
				j++
			}
		}
		return pre
	}

	var (
		getnum int64=2
		//res=make([]int64,0)
	)
	if uiDataInput<2 {
		return
	}
	for uiDataInput!=1{
		if uiDataInput % getnum == 0 {
			uiDataInput = uiDataInput / getnum
			fmt.Printf("%d ",getnum)
			//res = append(res, getnum)
		}else {
			getnum=nextnum(getnum)
			//logg.Printf("next is %d\n",getnum)

		}
	}
	//fmt.Printf("res is %v\n",res)
	return
}

//得到pre后面的第一个质数
func nextsu(pre int64) int64 {
	if pre==2 {
		return 3
	}
	var j int64=3
	pre+=2
	for j < pre / 2+1 {
		if pre%j==0 {
			pre+=2
			j=2
		} else {
			j++
		}
	}
	logg.Printf("next is %d\n",pre)
	return pre
}

func GetResult(uiDataInput int64) {
	if uiDataInput<2 {
		return
	}
	var divisor int64=2
	max:=uiDataInput/2
	for uiDataInput!=1 {
		if uiDataInput%divisor==0 {
			uiDataInput=uiDataInput/divisor
			fmt.Printf("%d ",divisor)
		} else {
			if divisor>max {
				fmt.Printf("%d ",uiDataInput)
				return
			}
			divisor++
		}
	}
	return
}

//写出一个程序，接受一个正浮点数值，输出该数值的近似整数值。
// 如果小数点后数值大于等于5,向上取整；小于5，则向下取整
func Approximation(num float32) int {
	//var res int
	//
	//res=int(num)
	//logg.Printf("res is %d,fr is %v\v",res,float32(res))
	//if num-float32(res)<0.5 {
	//	return res
	//} else {
	//	return res+1
	//}
	return int(num+0.5)
}

//数据表记录包含表索引和数值，请对表索引相同的记录进行合并，
//即将相同索引的数值进行求和运算，输出按照key值升序进行输出。
func mergeRepeat()  {
	var (
		count int
		key ,value int
	)
	fmt.Scanf("%d",&count)
	vmap:=make(map[int]int)
	ktmp:=make([]int,0)
	for i:=0;i<count ;i++  {
		fmt.Scanf("%d %d",&key,&value)
		if oldv,ok:=vmap[key];ok {
			vmap[key]=oldv+value
		} else {
			vmap[key]=value
			ktmp=append(ktmp,key)
		}

	}

	sort.Ints(ktmp)
	for _,v:=range ktmp {
		fmt.Printf("%d %d ",v,vmap[v])
	}
}

//输入一个int型整数，按照从右向左的阅读顺序，返回一个不含重复数字的新的整数
func ReverseIntNoRepeat(num int) int {
	if num==0 {
		return 0
	}
	var (
		i int
		str,pre,ress string
	)
	pre=strconv.Itoa(num)

	for i=0;i<len(pre);i++ {
		str=str+pre[len(pre)-i-1:len(pre)-i]
	}

	vmap:=make(map[byte]bool)
	for i := 0; i < len(str);i++ {
		if k, v := vmap[str[i]]; v&&k {
		}else {
			ress=ress+str[i:i+1]
			vmap[str[i]]=true
			//logg.Printf("add i is %d,v %c\n",i,str[i])
		}
	}

	res,err:=strconv.Atoi(ress)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return res
}

//编写一个函数，计算字符串中含有的不同字符的个数。
// 字符在ACSII码范围内(0~127)。不在范围内的不作统计。
func CountDiffLetter()  {
	var (
		b []byte
		//err error
		count int

	)
	vmap:=make(map[byte]bool)

	//InputReader:=bufio.NewReader(os.Stdin)
	//b,err=InputReader.ReadBytes('\n')
	//for err != nil {
	//	fmt.Println(err)
	//}
	fmt.Scanf("%s",&b)

	for i := 0; i < len(b); i++ {

		if b[i]>127 || b[i]<0{
			continue
		} else {
			if v, ok := vmap[b[i]]; v && ok {
				continue
			} else {
				vmap[b[i]]=true
				count++
			}
		}
	}
	fmt.Println(count)
}

//输入一个整数，将这个整数以字符串的形式逆序输出
//程序不考虑负数的情况，若数字含有0，则逆序形式也含有0，如输入为100，则输出为001
func GetReverseString() {
	var num int
	fmt.Scanf("%d",&num)
	s:=strconv.Itoa(num)
	end:=len(s)-1
	for i:=end;i>=0	 ;i--  {
		fmt.Printf("%c",s[i])
	}
}

func ReverseString() {
	var str string
	fmt.Scanf("%s",&str)
	end:=len(str)-1
	for i:=end-1;i>=0;i-- {
		fmt.Printf("%c",str[i])
	}
}

//将一个英文语句以单词为单位逆序排放。例如“I am a boy”，逆序排放后为“boy a am I”
//所有单词之间用一个空格隔开，语句中除了英文字母外，不再包含其他字符
func ReverseWords()  {
	//myl:=list.New()
	var str string
	inputReader:=bufio.NewReader(os.Stdin)
	str,err:=inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	end:=len(str)
	str=str[:end-1]
	res:=strings.Split(str," ")

	for i := len(res) - 1; i > 0; i-- {
		fmt.Printf("%s ",res[i])
	}
	fmt.Printf("%s",res[0])
	//fmt.Printf("res is %v,len is %v\n",res,len(res))
}
//输入第一行为一个正整数n(1≤n≤1000),
// 下面n行为n个字符串(字符串长度≤100),字符串中只含有大小写字母。
//给定n个字符串，请对n个字符串按照字典序排列
func StringsSort()  {
	var (
		num int
		str string
	)
	s:=make([]string,0)
	fmt.Scanf("%d",&num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%s",&str)
		s=append(s,str)
	}
	fmt.Printf("all str is %v\n",s)
	sort.Strings(s)
	//fmt.Printf("aflte str is %v\n",s)
	for i:=0;i<len(s);i++  {
		fmt.Println(s[i])
	}
}

//输入一个int型的正整数，计算出该int型数据在内存中存储时1的个数。
func Count1nums()  {
	const n int =1
	var (
		num,count int
	)
	fmt.Scanf("%d",&num)

	for num>0 {
		if num & n==1 {
			count++
		}
		num=num>>1
	}
	fmt.Printf("%d\n",count)
}