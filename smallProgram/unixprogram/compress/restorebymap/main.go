package main

import (
	"flag"
	"learntogo/smallProgram/unixprogram/compress/dedep"
	"log"
	"time"
)

func main() {
	var (
		msFile=flag.String("ms","","load ms file")
		comfile=flag.String("file","","common file")
		dst=flag.String("dst","","copy dst file")
	)
	flag.Parse()
	allmap,err:=dedep.GetValue(*comfile)
	if err!=nil {
		log.Println(err)
		return
	}
	ms,err:=dedep.LoadMs(*msFile)
	if err!=nil {
		log.Println(err)
		return
	}
	s:=time.Now()
	err=dedep.GenerFile(ms,allmap,*dst)
	if err!=nil {
		log.Println(err)
		return
	}
	log.Printf("cost time %v\n",time.Since(s))
}
