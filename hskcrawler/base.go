package main

import (
	"log"
	"os"
	"net/http"
	"fmt"
	"github.com/henrylee2cn/pholcus/common/goquery"
	"strconv"
	"regexp"
	"math"
)

var logg=log.New(os.Stdout,"log:", log.Lshortfile)

func main() {

	total()
	return

	var (
		name string="工作"
		url string="http://yuliaoku.hanyu123.cn/index.php/hsk/index/hsk?name="+name+"&c=&level=0&page=1"
	)
	//resp,err:=http.Get("http://yuliaoku.hanyu123.cn/index.php/hsk/index/hsk?name=%E5%B7%A5%E4%BD%9C&c=&level=0&page=1")
	resp,err:=http.Get(url)
	if err!=nil {
		logg.Fatalln(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode!=200 {
		logg.Printf("resp status is %v\n",resp.StatusCode)
		return
	}
	doc,err:=goquery.NewDocumentFromReader(resp.Body)
	if err!=nil {
		logg.Println(err)
		return
	}

	//fmt.Println(doc.Find("div.row").Find("table").Find("tbody").Text())
	logg.Println("11111")
	items:=doc.Find("div.row").Find("table").Find("tbody").Find("tr")
	items.Each(func(index int,sel *goquery.Selection) {
		if index==0 {
			return
		}
		a:=sel.Find("td")
		fmt.Printf("a:index is %v,tdfirst is %v,second is %v,third is %v\n",index,a.First().Text(),a.First().Next().Text(),a.First().Next().Next().Text())
		//a.Each(func(index int,sel *goquery.Selection) {
		//	fmt.Printf("index is %v,context is %v\n",index,sel.Text())
		//})
		//fmt.Printf("a index is %v,\n content is %v\n",index,a.Text())
	})

	logg.Println("all print")

	//fmt.Println(fsel.Text())

	logg.Println("each end")
	return

	lastpage:=doc.Find("ul.pagination").Find("li:nth-last-child(2)").Text()

	lastnum,err:=strconv.Atoi(lastpage)
	if err!=nil {
		logg.Println(err)
		return
	}

	fmt.Printf("last node is %v\n",lastnum)
	logg.Println("finish")

	return

}

func total()  {
	resp,err:=http.Get("http://www.baidu.com/s?ie=utf-8&nojc=1&wd=s7&rn=50")
	if err!=nil {
		logg.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode!=200 {
		logg.Printf("resp status is %v\n",resp.StatusCode)
		return
	}
	query,err:=goquery.NewDocumentFromReader(resp.Body)
	if err!=nil {
		logg.Println(err)
		return
	}
	total1 := query.Find(".nums").Text()
	logg.Printf("total1 is %v\n",total1)
	re, _ := regexp.Compile(`[\D]*`)
	total1 = re.ReplaceAllString(total1, "")
	total2, _ := strconv.Atoi(total1)
	logg.Printf("total2 is %v\n",total2)
	total := int(math.Ceil(float64(total2) / 50))
	logg.Printf("total is %v\n",total)

}
