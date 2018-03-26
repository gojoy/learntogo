package dedep

import (
	"bufio"
	"crypto/md5"
	"encoding/gob"
	"io"
	"log"
	"os"
)

type tb map[[16]byte]uint64

func GetTB(path string) (tb, error) {
	var (
		err                error
		dict               = make(tb)
		buf                = make([]byte, blocksize)
		offset, count, all uint64
		num                [16]byte
	)
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer f.Close()
	bf := bufio.NewReader(f)
	for {
		_, err = bf.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Println(err)
			}
		}
		offset = offset + blocksize
		all++
		num = md5.Sum(buf)
		if _, ok := dict[num]; ok {
			count++
		} else {
			dict[num] = offset
		}
		if all%100 == 0 {
			log.Printf("offset is %v,bufferd is %v\n", offset, bf.Buffered())
		}
	}
	log.Printf("len is %v,all is %v\n", len(dict), all)
	return dict, nil
}

func DumpToFile(dict tb, path string) error {
	var (
		err error
	)
	f, err := os.Create(path)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	if err = enc.Encode(dict); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
