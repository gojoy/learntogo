package main

import (
	"log"
	"os"
	"time"
	"os/exec"
	"fmt"
)

var (
	glog=log.New(os.Stderr,"mytime",log.Lshortfile)
)
func main() {
	if len(os.Args)<2 {
		glog.Println("usage:mytime command...")
	}
	myargs:=os.Args[1:]
	startTime:=time.Now()
	execrun(myargs)
	fmt.Printf("time %v\n",time.Since(startTime))
}

func execrun(args []string)  {
	var mycmd *exec.Cmd
	if len(args)==1 {
		mycmd=exec.Command(args[0],)
	} else {
		mycmd=exec.Command(args[0],args[1:]...)
	}
	err:=mycmd.Run()
	if err!=nil {
		glog.Println(err)
		os.Exit(1)
	}
	return
}
