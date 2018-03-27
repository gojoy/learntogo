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
	ms,err:=dedep.LoadMs(*msFile)
	if err!=nil {
		log.Println(err)
		return
	}
	s:=time.Now()
	err=dedep.GenerFileByOffset(ms,*dst,*comfile)
	if err!=nil {
		log.Println(err)
		return
	}

	log.Printf("load ms:%v,common file:%v,make dst file:%v\n",*msFile,*comfile,*dst)
	log.Printf("restore file cost:%v\n",time.Since(s))
}
