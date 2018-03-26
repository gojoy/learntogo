package main

import (
	"bufio"
	"crypto/md5"
	"io"
	"log"
	"os"
)

func GetPagesbio(path string) (map[uint16]bool, error) {
	var (
		dic   = make(map[uint16]bool)
		count uint64
		num   uint16
		all   uint64
		//b uint16
	)
	buf := make([]byte, 2)
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	bf := bufio.NewReader(f)
	for {
		//_,err=io.CopyN(buf,f,2)
		_, err = bf.Read(buf)
		//_,err=f.Read(buf)
		if err != nil {
			log.Println(err)
			break
			//return err
		}
		all++
		//log.Printf("buf is %x,%x\n",buf[0],buf[1])
		//log.Printf("0 is %v,1 is %v",uint16(buf[0]),uint16(buf[1]))
		//log.Printf("sum is %v\n",uint16(buf[0])<<8+uint16(buf[1]))
		num = uint16(buf[0])<<8 + uint16(buf[1])
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

func getMd5Page(path string) (map[[16]byte]bool, error) {
	var (
		err         error
		dict        = make(map[[16]byte]bool)
		counts, all uint
	)
	buf := make([]byte, 4096)
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	bf := bufio.NewReader(f)
	for {
		_, err = bf.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		r := md5.Sum(buf)
		if _, ok := dict[r]; ok {
			counts++
		} else {
			dict[r] = true
		}
		all++
	}
	log.Printf("all is %v,count is %v,len is %v\n", all, counts, len(dict))
	return dict, nil
}

func getSameMd5(path string, dict map[[16]byte]bool) error {
	var (
		counts, all uint
	)
	buf := make([]byte, 4096)
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return err
	}
	bf := bufio.NewReader(f)
	for {
		_, err = bf.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Println(err)
				return err
			}
		}
		m := md5.Sum(buf)
		if _, ok := dict[m]; ok {
			counts++
		}
		all++
	}
	log.Printf("same counts is %v,all is %v\n", counts, all)
	return nil
}
