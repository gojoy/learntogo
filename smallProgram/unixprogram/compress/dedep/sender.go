package dedep

import (
	"log"
	"os"
	"bufio"
	"io"
	"crypto/md5"
	"encoding/gob"
)

var (
	FirstMd5 [16]byte
	FirstOffset int64
)
const blocksize  =4096

type sdata struct {
	Soffset int64
	Data    []byte
}

type ddata struct {
	Md5num           [16]byte
	Soffset, Doffset int64
}

type message struct {
	Isd bool
	S sdata
	D ddata
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

//去重，得到path1-path的去重文件
func DoDep(path string,dict tb) ([]message,error)  {

	var (
		err error
		res=make([]message,0)
		buf=make([]byte,blocksize,blocksize)
		sum [16]byte
		offset int64
		count uint64
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
		if n!=blocksize {
			log.Printf("n is %v\n",n)
			return nil,err
		}
		sum=md5.Sum(buf)

		if v,ok:=dict[sum];ok {

			dd:=ddata{
				Md5num:sum,
				Soffset:offset,
				Doffset:v,
			}
			res=append(res,message{D:dd,Isd:true})
			//log.Printf("num %d,md5 is %v,buf is %v\n",all,sum,buf)
			//FirstOffset=v
			//FirstMd5=sum
		} else {
			senddata:=make([]byte,blocksize,blocksize)
			copyn:=copy(senddata,buf)
			if copyn!=blocksize {
				log.Printf("copy error:from %v\n to\n%v\n",buf,senddata)
				panic("")
			}
			sd:=sdata{
				Soffset:offset,
				Data:senddata,
			}
			res=append(res,message{S:sd,Isd:false})
			//log.Printf("in %v,num is %v,offset is %v,data is %x\n",path,all,offset,buf)
			//FirstOffset=offset
			//return nil,nil
		}
		count++
		offset=offset+blocksize
		if count%100==0 {
			//log.Println(res[len(res)-1])
		}
	}

	log.Printf("message len is %v\n",len(res))
	log.Printf("last message is %v\n",res[len(res)-1])
	return res,nil
}

func SaveMessageBin(ms []message,path string) error {
	var (
		err error
		//path="/opt/mes.gob"
	)
	f,err:=os.Create(path)
	if err!=nil {
		log.Println(err)
		return err
	}
	defer f.Close()
	enc:=gob.NewEncoder(f)
	if err=enc.Encode(ms);err!=nil {
		log.Println(err)
		return err
	}
	//log.Printf("len :%v,v:%v\n",len(ms),ms[0])
	return nil
}