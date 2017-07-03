package huawei

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"

	"container/list"
)


//开发一个坐标计算工具， A表示向左移动，D表示向右移动，W表示向上移动，S表示向下移动。从（0,0）点开始移动，从输入字符串里面读取一些坐标，并将最终输入结果输出到输出文件里面。
//
//输入：
//
//合法坐标为A(或者D或者W或者S) + 数字（两位以内）
//
//坐标之间以;分隔。
//
//非法坐标点需要进行丢弃。如AA10;  A1A;  $%$;  YAD; 等。
//
//下面是一个简单的例子 如：
//
//A10;S20;W10;D30;X;A1A;B10A11;;A10;
//
//处理过程：
//
//起点（0,0）
//
//+   A10   =  （-10,0）
//
//+   S20   =  (-10,-20)
//
//+   W10  =  (-10,-10)
//
//+   D30  =  (20,-10)
//
//+   x    =  无效
//
//+   A1A   =  无效
//
//+   B10A11   =  无效
//
//+  一个空 不影响
//
//+   A10  =  (10,-10)
//
//
//
//结果 （10， -10）
func GetCoordinate() {
	const Operate string = "ASWD"
	const step string="0123456789"
	var (
		num int
	)
	type position struct {
		x int
		y int
	}
	p:=position{
		x:0,
		y:0,
	}
	reader:=bufio.NewReader(os.Stdin)

	s,err:=reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("input is %v\n",s)
	input:=strings.Split(s,";")

	for i := 0; i < len(input); i++ {
		if len(input[i])==0 {
			continue
		}
		if num,err=strconv.Atoi(input[i][1:]);err!=nil {
			//fmt.Printf("atoi err is %v\n",err)
			continue
			}
		switch input[i][0] {
		case 'A':
			p.x-=num
		case 'D':
			p.x+=num
		case 'W':
			p.y+=num
		case 'S':
			p.y-=num
		default:
			continue
		}
	}
	fmt.Printf("%d,%d\n",p.x,p.y)

	return
}



//开发一个简单错误记录功能小模块，能够记录出错的代码所在的文件名称和行号。
//
//处理：
//
//1、 记录最多8条错误记录，循环记录，对相同的错误记录（净文件名称和行号完全匹配）只记录一条，错误计数增加；
//
//2、 超过16个字符的文件名称，只记录文件的最后有效16个字符；
//
//3、 输入的文件可能带路径，记录文件名称不能带路径。
//
//
//输入描述:
//一行或多行字符串。每行包括带路径文件名称，行号，以空格隔开。
//
//
//输出描述:
//将所有的记录统计并将结果输出，格式：文件名 代码行数 数目，一个空格隔开，如：
//示例1
//输入
//
//E:\V1R2\product\fpgadrive.c   1325
//输出
//
//fpgadrive.c 1325 1
func CodeLoger() {
	type putlog struct {
		name string
		line int
		num int
	}

	var (
		str string
		path string
		line int
		err error
	)
	msgList:=list.New()

	dealMsg:= func(s string) {
		if len(s) == 0 {
			return
		}
		sl:=strings.Split(s," ")

		allpath:=strings.SplitAfter(sl[0],"\\")
		path=allpath[len(allpath)-1]
		if len(path)>16 {
			path=path[len(path)-16:]
		}
		line,err=strconv.Atoi(sl[len(sl)-1])
		if err != nil {
			fmt.Println(err)
			return
		}
		if msgList.Len()==0 {
			msgList.PushFront(putlog{
				name:path,
				line:line,
				num:1,
			})
		} else {
			for e:=msgList.Front();e!=nil;e=e.Next() {

				if vl,ok:=e.Value.(putlog);ok &&vl.line==line && vl.name==path {
					vl.num++
					return
				}
			}
			msgList.PushFront(putlog{
				name:path,
				line:line,
				num:1,
			})
			if msgList.Len()>8 {
				msgList.Remove(msgList.Back())
			}
		}

	}
	//di,err:=ioutil.ReadAll(os.Stdin)
	//if err != nil {
	//	println(err)
	//}
	//fmt.Printf("%s\n",di)
	scaner:=bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		str=scaner.Text()
		dealMsg(str)
	}
	if err := scaner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for e:=msgList.Front();e!=nil;e=e.Next() {
		vl,ok:=e.Value.(putlog)
		if ok {
			fmt.Printf("%s %d %d\n",vl.name,vl.line,vl.num)
		}
	}


}