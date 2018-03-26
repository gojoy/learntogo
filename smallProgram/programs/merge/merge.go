package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	src  = flag.String("src", "/run/migration/upper/.", "copy src")
	dst  = flag.String("dst", "/var/lib/docker/workfile/overlaytest/lazy", "copy dst")
	merg = flag.String("merge", "/var/lib/docker/workfile/overlaytest/merged", "merge dir")
	err  error
)

/*
func main() {
	err=os.Chdir("/run/migration/")
	log.Println(err)
	if err!=nil {
		return
	}

	log.SetFlags(log.Lshortfile | log.Ltime)
	flag.Parse()
	fmt.Printf("src is %v,dst is %v,merged is %v\n", *src, *dst, *merg)

	if err = MergeDir(*src, *dst, *merg); err != nil {
		log.Println(err)
		return
	}
	return
}
*/

func MergeDir(upper, lazy, mergedir string) error {

	var (
		err error
	)

	args := []string{"-rf"}
	args = append(args, upper, lazy)
	cmd := exec.Command("cp", args...)
	log.Printf("cmd is %v\n", cmd.Args)
	for _, v := range cmd.Args {
		fmt.Printf("args is %v.\n", v)
	}

	buf, err := cmd.CombinedOutput()
	log.Printf("err is %v,out is %v\n", err, string(buf))
	if err != nil {
		log.Printf("err is %v,out is %v\n", err, string(buf))
		return err
	}
	//if err = cmd.Run(); err != nil {
	//	log.Println(err)
	//	return err
	//}
	if err = os.RemoveAll(mergedir); err != nil {
		log.Println(err)
		return err
	}

	if err = os.Rename(lazy, mergedir); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
