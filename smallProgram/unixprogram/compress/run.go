package main

import (
	"flag"
	"learntogo/smallProgram/unixprogram/compress/dedep"
	"log"
)

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}

const p = "/opt/dict.bin"

func main() {
	var (
		err error
	)
	log.Println("start")

	path := flag.String("dir", "", "common file")
	path1 := flag.String("dir1", "", "copy src file")
	path2:=flag.String("tar","","copy target file")
	flag.Parse()
	log.Println(*path, *path1)
	//path: common file
	log.Println("start get maptable")
	dict1, err := dedep.GetTB(*path)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("save dict")
	err=dedep.DumpToFile(dict1,"/opt/dict.gob")
	if err!=nil {
		log.Println(err)
		return
	}
	log.Println("start make diff file")
	//path1: rsync file
	//ms:diff file
	ms,err:=dedep.DoDep(*path1,dict1)
	if err!=nil {
		log.Println(err)
		return
	}
	log.Println("save message")
	err=dedep.SaveBin(ms,"/opt/ms.gob")
	if err!=nil {
		log.Println(err)
		return
	}
	//log.Printf("fisrs is %v\nfirst offset is %v\n",dedep.FirstMd5,dedep.FirstOffset)


	log.Println("start gener file")
	err=dedep.GenerFileByOffset(ms,*path2,*path)
	if err!=nil {
		log.Println(err)
		return
	}
	log.Println("done!")

}
