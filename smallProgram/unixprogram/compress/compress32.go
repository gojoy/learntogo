package main

import (
	"log"
	"os"
)

func GetPages32(path string) (map[uint32]bool, error) {
	var (
		dic   = make(map[uint32]bool)
		count uint64
		num   uint32
		all   uint64
		//b uint16
	)
	buf := make([]byte, 4)

	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for {
		//_,err=io.CopyN(buf,f,2)
		_, err = f.Read(buf)
		if err != nil {
			log.Println(err)
			break
			//return err
		}
		all++
		//log.Printf("buf is %x\n",buf)
		//log.Printf("0 is %v,1 is %v",uint16(buf[0]),uint16(buf[1]))
		//log.Printf("sum is %v\n",uint16(buf[0])<<8+uint16(buf[1]))
		num = uint32(buf[0])<<24 + uint32(buf[1])<<16 + uint32(buf[2])<<8 + uint32(buf[3])
		if _, ok := dic[num]; ok {
			count++
			//log.Println(num)
		} else {
			dic[num] = true
		}
		//buf.Reset()
	}
	log.Printf("count is %v,len map is %v,all is %v\n", count, len(dic), all)
	return dic, nil
}

func GetSame32(path string, dict map[uint32]bool) error {
	var (
		nums         uint32
		counts, same uint64
	)
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return err
	}
	buf := make([]byte, 4)
	for {
		last, err := f.Read(buf)
		if err != nil {
			log.Printf("last is: %v,err is %v\n", last, err)
			break
		}
		counts++
		nums = uint32(buf[0])<<24 + uint32(buf[1])<<16 + uint32(buf[2])<<8 + uint32(buf[3])
		if _, ok := dict[nums]; ok {
			same++
		}
	}
	log.Printf("same is %v,count is %v\n", same, counts)
	return nil
}
