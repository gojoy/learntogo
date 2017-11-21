package main

import (
	"log"
	"os"
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/henrylee2cn/pholcus/common/goquery"
	"strconv"
)

var logg=log.New(os.Stdout,"log:", log.Lshortfile)

func main() {
	resp,err:=http.Get("http://yuliaoku.hanyu123.cn/index.php/hsk/index/hsk?name=%E5%B7%A5%E4%BD%9C&c=&level=0&page=1")
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
		a:=sel.Find("td")
		a.Each(func(index int,sel *goquery.Selection) {
			fmt.Printf("index is %v,context is %v\n",index,sel.Text())
		})
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

	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil {
		logg.Println(err)
		return
	}
	fmt.Println(string(body))

}
