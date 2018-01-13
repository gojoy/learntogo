package qunar

import (
	"bufio"
	"os"
	"fmt"
	"sort"
	"strconv"
)

type bybyte []byte

func (a bybyte)Len() int  {
	return len(a)
}
func (a bybyte)Swap(i, j int)  {
	a[i], a[j] = a[j], a[i]
}

func (a bybyte)Less(i,j int) bool  {
	return a[i]<a[j]
}

func Q1()  {
	var (
		k int
		str,str1 string
		err error
	)
	bs:=make([]byte,0)
	input:=bufio.NewReader(os.Stdin)
	str,err=input.ReadString('\n')
	//if err!=nil {
	//	fmt.Println(err)
	//	return
	//}
	str1,err=input.ReadString('\n')
	if str[len(str)-1]=='\n' {
		str=str[:len(str)-1]
	}
	if str1[len(str1)-1]=='\n' {
		str1=str1[:len(str1)-1]
	}
	//fmt.Printf("str is %v,str1 is %v\n",str,str1)
	k,err=strconv.Atoi(str1)
	if err!=nil {
		fmt.Println(err)
		return
	}
	for i:=0;i<len(str);i++ {
		if str[i]!=32 {
			bs=append(bs,str[i])
		}
	}
	sort.Sort(bybyte(bs))
	//fmt.Printf("k is %v,str is %v\n",k,string(bs))
	fmt.Println(string(bs[:k]))
	l:=len(bs)-1
	for i:=0;i<k;i++ {
		fmt.Printf("%c",bs[l-i])
	}
	fmt.Printf("\n")
}

