package main

import (
	"learntogo/gopl/ch4/github"
	"os"
	"log"
	"fmt"
	"encoding/json"
)

func main() {


	result,err:=github.SearchIssues(os.Args[1:])
	if err!=nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d issues:\n",result.TotalCount)
	for _,item:=range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	data,err:=json.MarshalIndent(result,"","	")
	//data,err:=json.MarshalIndent(*result,"","	")
	if err!=nil {
		log.Fatalln(err)
	}

	fp,err:=os.OpenFile("github.txt",os.O_CREATE|os.O_RDWR|os.O_TRUNC,0666)
	if err!=nil {
		log.Fatalln(err)
	}
	defer fp.Close()

	_,err=fp.Write(data)
	if err!=nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(fp,"\nxml add by ch4!\n")

	fmt.Printf("data is %v\n",string(data))
}
