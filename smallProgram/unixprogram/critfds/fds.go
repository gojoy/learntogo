package main

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
	"strings"
	"time"
	//"path/filepath"
	"os"
)

type Volumes struct {
	src, dst string
}

// 0: local src copy dir  1: local relative file path
//[/var/lib//docker../data t1/t1.ibd]
type allpath [2]string

// each volumes dir,need to sync files
type volpath []allpath

var (
	path = "/run/migration/d5c02022a630311f4451dc89aec4257192751b944a7846fe8f9f51868ca93b08/fullDump"

	v = []Volumes{struct{ src, dst string }{src: "/var/lib/docker/workfile/vols/data", dst: "/var/lib/mysql/"}}
)

func main() {
	var (
		err   error
		files []string
		ip    = "192.168.18.128"
		r     = "/tmp/files"
	)
	s := time.Now()
	log.SetFlags(log.Lshortfile | log.Ltime)
	if files, err = getFiles(path); err != nil {
		log.Println(err)
	}

	p, err := syncNeedFiles(files, v)
	if err != nil {
		log.Println(err)
		return
	}
	//log.Println(p)
	if err = fastCopy(p, ip, r); err != nil {
		log.Println(err)
		return
	}

	log.Printf("cost %v\n", time.Since(s))
	return
}

func getFiles(path string) ([]string, error) {
	var (
		err error
		res = make([]string, 0)
	)

	args := []string{"x"}
	args = append(args, path, "fds")
	cmd := exec.Command("crit", args...)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return res, err
	}
	//log.Printf("out is %v\n",string(out))

	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		txt := s.Text()
		//log.Println(s.Text())
		sp := strings.Split(txt, ":")
		//log.Printf("sp is %v,len is %v\n",sp,len(sp))
		if len(sp) == 2 {
			if sp[1][1] == '/' {
				res = append(res, sp[1][1:])
			}
		}
	}
	res = res[:len(res)-2]
	//log.Printf("res is %v,len is %v\n",res,len(res))
	return res, nil
}

func syncNeedFiles(files []string, vol []Volumes) ([]volpath, error) {
	var (
		err   error
		realp = make([]allpath, 0)
		res   = make([]volpath, 0)
	)
	if len(vol) == 0 {
		log.Println("vol NULL")
		return res, nil
	}

	for _, v := range vol {
		for i := 0; i < len(files); i++ {
			right := strings.TrimPrefix(files[i], v.dst)
			if len(right) != len(files[i]) {
				//log.Printf("path is %v,right is %v\n",files[i],right)
				//l:=filepath.Join(v.src,right)
				//log.Printf("copy file is %v\n",l)
				realp = append(realp, allpath{v.src, right})
			}
		}
		res = append(res, realp)
	}
	//log.Printf("len is %v\n",len(res))
	return res, err
}

func fastCopy(files []volpath, ip string, remote string) error {
	var (
		err error
	)
	if len(files) == 0 {
		return err
	}
	if err = os.Chdir(v[0].src); err != nil {
		log.Println(err)
		return err
	}
	//log.Printf("f is %v\n",files)
	for _, v := range files {
		if err = os.Chdir(v[0][0]); err != nil {
			log.Println(err)
		}
		for _, v1 := range v {
			if err = RemoteCopyDirRsync(v1[1], remote, ip); err != nil {
				log.Println(err)
			}
		}
		//log.Println(v)
		//rpath:=filepath.Join(remote,v[1])
		//log.Println(rpath)
		//if err=RemoteCopyDirRsync(v[1],remote,ip);err!=nil {
		//	log.Println(err)
		//	return err
		//}
	}
	return nil
}

func RemoteCopyDirRsync(local, remote string, ip string) error {

	var (
		err error
	)
	//if local[len(local)-1] != '/' {
	//	local = local + "/"
	//}
	//if remote[len(remote)-1] != '/' {
	//	remote = remote + "/"
	//}

	args := append([]string{"-azR"}, local, "root@"+ip+":"+remote)
	//log.Printf("l is %v,r is %v,args is %v\n",local,remote,args)

	cmd := exec.Command("rsync", args...)
	log.Println(cmd.Args)
	//log.Printf("cmd is %v\n",cmd)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("rsync error:%v,out:%v\n", err, string(out))
		log.Printf("cmd is %v\n", cmd.Args)
	}
	return err
}
