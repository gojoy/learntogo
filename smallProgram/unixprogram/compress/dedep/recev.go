package dedep

import (
	"bufio"
	"crypto/md5"
	"encoding/gob"
	"io"
	"log"
	"os"
)

//对应md5与当前文件的偏移量
type tb map[[16]byte]int64

type db map[[16]byte][]byte

func GetTB(path string) (tb, error) {
	var (
		err                error
		dict               = make(tb)
		buf                = make([]byte, blocksize,blocksize)
		offset int64
		count, all uint64
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

		all++
		num = md5.Sum(buf)
		if _, ok := dict[num]; ok {
			count++
		} else {
			dict[num] = offset
		}
		offset = offset + blocksize
		//if all%100 == 0 {
		//	log.Printf("offset is %v,bufferd is %v\n", offset, bf.Buffered())
		//}
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

func GetValue(path string)(db,error) {
	var (
		err error
		kv=make(db)
		buf=make([]byte,blocksize)
		sum [16]byte
		counts uint64
		n int
	)
	f,err:=os.Open(path)
	if err!=nil {
		log.Println(err)
		return nil,err
	}
	defer f.Close()
	bf:=bufio.NewReader(f)
	for  {
		n,err=bf.Read(buf)
		if err!=nil {
			if err==io.EOF {
				break
			} else {
				log.Println(err)
				return nil,err
			}
		}
		//log.Printf("n is %v\n",n)
		sum=md5.Sum(buf)
		if _,ok:=kv[sum];ok {
			counts++
		} else {
			savedata:=make([]byte,blocksize,blocksize)
			copyn:=copy(savedata,buf)
			if copyn!=blocksize {
				log.Println("copy error:\n",buf)
				panic("")
			}
			kv[sum]=savedata
		}
	}
	log.Printf("last n is %v\n",n)
	return kv,nil
}

func GenerFile(ms []message,kv db,path string ) error  {
	var (
		err error
	)
	if len(ms)==0 {
		log.Println("nil message")
		return err
	}
	f,err:=os.Create(path)
	if err!=nil {
		log.Println(err)
		return err
	}
	bf:=bufio.NewWriter(f)
	defer f.Close()
	for i:=0;i<len(ms);i++ {

		if ms[i].Isd {
			data,ok:=kv[ms[i].D.Md5num]
			if !ok {
				log.Printf("not match:%v\n",ms[i].D.Md5num)
				panic("false\n")
			}
			_,err=bf.Write(data)
			if err!=nil {
				log.Println(err)
				return err
			}
		} else {
			_,err=bf.Write(ms[i].S.Data)
			if err!=nil {
				log.Println(err)
				return err
			}
		}
	}
	return nil
}



func GenerFileByOffset(ms []message, path,oldpath string) error {
	//buf:=make([]byte,blocksize)
	if len(ms)==0 {
		panic("nil ms\n")
	}
	f,err:=os.Create(path)
	if err!=nil {
		log.Println(err)
		return err
	}
	defer f.Close()
	old,err:=os.Open(oldpath)
	if err!=nil {
		log.Println(err)
		return err
	}
	defer old.Close()

	for i:=0;i<len(ms);i++ {
		if ms[i].Isd {
			_,err=old.Seek(ms[i].D.Doffset,0)
			if err!=nil {
				log.Println(err)
				return err
			}
			_,err=io.CopyN(f,old,blocksize)
			if err!=nil {
				log.Println(err)
				return err
			}
		} else {
			_,err=f.Write(ms[i].S.Data)
			if err!=nil {
				log.Println(err)
				return err
			}
		}
	}
	return nil
}

func LoadMs(path string) ([]message, error) {
	var (
		err error
		ms=make([]message,0)
	)
	f,err:=os.Open(path)
	if err!=nil {
		log.Println(err)
		return nil,err
	}
	defer f.Close()

	dec:=gob.NewDecoder(f)
	err=dec.Decode(&ms)
	if err!=nil {
		log.Println(err)
		return nil,err
	}
	return ms,nil
}

