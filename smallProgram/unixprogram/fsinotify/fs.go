package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"learntogo/smallProgram/unixprogram/fsinotify/inotifyTree"
	"log"
	"os"
)

var (
	glog = log.New(os.Stderr, "", log.Lshortfile)
	dir  = flag.String("dir", "/tmp", "add watch dir")
	err  error
)

func main() {

	w, err := fsnotify.NewWatcher()
	if err != nil {
		glog.Fatalln(err)
	}

	flag.Parse()

	if flag.NArg() != 0 {
		glog.Fatalln("Usage:fsinotify -dir /tmp")
	}

	fmt.Printf("dir is %v\n", *dir)

	defer w.Close()
	if err = inotifyTree.AddWatchAll(*dir, w); err != nil {
		glog.Fatalln(err)
	}

	fmt.Println("start monitor")
	for {
		select {
		case event := <-w.Events:
			glog.Println(event)
			if event.Op&fsnotify.Create == fsnotify.Create {
				glog.Println("create file:", event.Name)
				info, err := os.Stat(event.Name)
				if err == nil && info.IsDir() {
					glog.Printf("update: add %v to list\n", event.Name)
					if err = w.Add(event.Name); err != nil {
						glog.Println(err)
					}
				}
			}
			if event.Op&fsnotify.Remove == fsnotify.Remove {
				glog.Println("remove file:", event.Name)
			}

		case e := <-w.Errors:
			glog.Fatalln(e)
			return
		}
	}
}
