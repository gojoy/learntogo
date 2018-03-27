package main

import (
	"flag"
	"learntogo/smallProgram/unixprogram/compress/dedep"
	"log"
	"time"
)

func main() {
	var (
		dictFile=flag.String("dict","","load dict file")
		file=flag.String("src","","copy src file")
		msFile=flag.String("ms","","message dump file")
	)
	flag.Parse()
	dict,err:=dedep.LoadDict(*dictFile)
	if err!=nil {
		log.Println(err)
		return
	}
	s:=time.Now()
	ms,err:=dedep.DoDep(*file,dict)
	if err!=nil {
		log.Println(err)
		return
	}
	log.Printf("make diff cost %v\n",time.Since(s))
	err=dedep.SaveMessageBin(ms,*msFile)
	if err!=nil {
		log.Println(err)
	}
	log.Printf("load from %v,copy %v,diff file %v\n",*dictFile,*file,*msFile)
	return
}
