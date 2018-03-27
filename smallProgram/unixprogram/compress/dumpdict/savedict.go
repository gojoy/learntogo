package main

import (
	"flag"
	"learntogo/smallProgram/unixprogram/compress/dedep"
	"log"
)

func main() {
	var (
		file=flag.String("file","","input common file")
		fil1=flag.String("save","","save dict to file")
	)
	flag.Parse()
	dict,err:=dedep.GetTB(*file)
	if err!=nil {
		log.Println(err)
		return
	}
	err=dedep.DumpToFile(dict,*fil1)
	if err!=nil {
		log.Println(err)
	}
	log.Printf("save %v to dict %v\n",*file,*fil1)
	return
}
