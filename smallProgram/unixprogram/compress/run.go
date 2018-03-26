package main

import (
	"flag"
	"learntogo/smallProgram/unixprogram/compress/dedep"
	"log"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

const p = "/opt/dict.bin"

func main() {
	var (
		err error
	)
	log.Println("start")
	//if err:=loadmap("/opt/dict.gob");err!=nil {
	//	log.Println(err)
	//}
	//return
	path := flag.String("dir", "", "read page file")
	path1 := flag.String("dir1", "", "comp img file")
	flag.Parse()
	log.Println(*path, *path1)
	dict1, err := dedep.GetTB(*path)
	if err != nil {
		log.Println(err)
		return
	}
	if err = dedep.DumpToFile(dict1, p); err != nil {
		log.Println(err)
		return
	}
	if _, err = dedep.LoadDict(p); err != nil {
		log.Println(err)
		return
	}
	res,err:=dedep.DoDep(*path1,dict1)
	if err!=nil {
		log.Println(err)
		return
	}
	if err=dedep.SaveJson(res);err!=nil {
		log.Println(err)
	}


	return
	dict, err := getMd5Page(*path)
	if err != nil {
		log.Println(err)
	}

	if err = getSameMd5(*path1, dict); err != nil {
		log.Println(err)
	}
	if err = save(dict); err != nil {
		log.Println(err)
	}
	return
	//d,err:=GetPages(*path)
	//if err!=nil {
	//	log.Println(err)
	//}
	//log.Printf("now comp %v\n",*path1)
	//return
	//err=GetSame(*path1,d)
	//if err!=nil {
	//	log.Println(err)
	//}
}
