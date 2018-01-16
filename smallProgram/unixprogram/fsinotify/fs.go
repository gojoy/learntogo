package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"flag"
	"fmt"
	"learntogo/smallProgram/unixprogram/fsinotify/inotifyTree"
)

var (
	glog=log.New(os.Stderr,"",log.Lshortfile)
	dir=flag.String("dir","/tmp","add watch dir")
	err error
)

func main() {

	w,err:=fsnotify.NewWatcher()
	if err!=nil {
		glog.Fatalln(err)
	}

	flag.Parse()


	if flag.NArg()!=0 {
		glog.Fatalln("Usage:fsinotify -dir /tmp")
	}

	fmt.Printf("dir is %v\n",*dir)

	defer w.Close()
	if err=w.Add(*dir);err!=nil {
		glog.Fatalln(err)
	}
	if err=inotifyTree.AddWatchAll(*dir,w);err!=nil {
		glog.Fatalln(err)
	}

	fmt.Println("start monitor")

	go inotifyTree.UpdateDirTree(w)

	for {
		select {
		case event := <-w.Events:
			glog.Println(event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				glog.Println("modified file:", event.Name)
			}
			if event.Op&fsnotify.Remove==fsnotify.Remove {
				glog.Println("remove file:",event.Name)
			}

		case e := <-w.Errors:
			glog.Fatalln(e)
			return
		}
	}


}
