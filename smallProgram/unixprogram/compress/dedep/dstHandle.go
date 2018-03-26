package dedep

import (
	"encoding/gob"
	"log"
	"os"
	"bufio"
	"io"
	"crypto/md5"
	"encoding/json"
)

const blocksize=4096

type sdata struct {
	soffset uint64
	data    []byte
}

type ddata struct {
	md5num           [16]byte
	soffset, doffset uint64
}

type message struct {
	isd bool
	s sdata
	d ddata
}

func LoadDict(path string) (tb, error) {
	var (
		err  error
		dict = make(tb)
	)
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer f.Close()
	doc := gob.NewDecoder(f)
	if err = doc.Decode(&dict); err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("len is %v\n", len(dict))
	return dict, nil
}

func DoDep(path string,dict tb) ([]message,error)  {

	var (
		err error
		res=make([]message,0)
		buf=make([]byte,blocksize)
		sum [16]byte
		offset uint64
		count uint64
	)
	f,err:=os.Open(path)
	if err!=nil {
		log.Println(err)
		return nil,err
	}
	defer f.Close()
	bf:=bufio.NewReader(f)

	for {
		_,err=bf.Read(buf)
		if err!=nil {
			if err==io.EOF {
				break
			} else {
				log.Println(err)
				return nil,err
			}
		}
		sum=md5.Sum(buf)
		offset=offset+blocksize

		if v,ok:=dict[sum];ok {
			dd:=ddata{
				md5num:sum,
				soffset:offset,
				doffset:v,
			}
			res=append(res,message{d:dd,isd:true})
		} else {
			sd:=sdata{
				soffset:offset,
				data:buf,
			}
			res=append(res,message{s:sd,isd:false})
		}
		count++
		if count%100==0 {
			//log.Println(res[len(res)-1])
		}
	}
	log.Printf("len is %v\n",len(res))
	return res,nil
}

func SaveJson(ms []message) error {
	var (
		err error
		path="/opt/mes.json"
	)
	f,err:=os.Create(path)
	if err!=nil {
		log.Println(err)
		return err
	}
	defer f.Close()
	enc:=json.NewEncoder(f)
	if err=enc.Encode(ms);err!=nil {
		log.Println(err)
		return err
	}
	log.Printf("len :%v,v:%v\n",len(ms),ms[0])
	return nil
}