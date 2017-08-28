package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var (
//		str string
//		err error
//		flag int
//		j int
//		i int
//	)
//	res:=make([]string,0)
//
//	fmt.Printf("ok 3\n")
//	_,err=fmt.Scanf("%s",&str)
//	if err!=nil {
//		return
//	}
//	for i=0;i<len(str);i++ {
//		if str[i]=='@' {
//			if flag==1 {
//				flag=0
//				j=i
//			}else {
//				flag=1
//			}
//		}
//		if(str[i]=='.') {
//			if flag==1 {
//				res=append(res,str[j:i+1])
//				flag=0
//				i++
//				j=i
//			}
//		}
//	}
//	handemail1(res)
//}
//
//func handemail1(s []string)  {
//	if len(s)==0 {
//		return
//	}
//	//fmt.Printf("handle\n")
//	for i:=0;i<len(s);i++ {
//		if len(s[i])>124 {
//			continue
//		}
//		//fmt.Printf("si is %s\n",s[i])
//		ptemail(s[i])
//	}
//}
//func ptemail(str string)  {
//	var (
//		countleft int
//		i int
//	)
//	if len(str)==0 {
//		return
//	}
//	for i=0;str[i]!='@';i++ {
//		countleft++
//	}
//	if countleft<3 {
//		return
//	}
//	if len(str)-i-1>119 {
//		return
//	}
//	//fmt.Printf("nihua is \n%s\n",str)
//	strsli:=[]byte(str)
//	strsli[i-1]='*'
//	strsli[i-2]='*'
//	strsli[i-3]='*'
//	fmt.Printf("%s\n",strsli)
//}
