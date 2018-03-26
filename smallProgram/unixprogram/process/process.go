package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var (
	glog = log.New(os.Stderr, "log ", log.Lshortfile)
)

func main() {

	if hostName, err := os.Hostname(); err != nil {
		glog.Fatalln(err)
	} else {
		fmt.Printf("hostname %v\n", hostName)
	}

	if dir, err := os.Getwd(); err != nil {
		glog.Fatalln(err)
	} else {
		fmt.Printf("pwd is %v\n", dir)
	}
	fmt.Printf("pid %v,\tppid %v\n", os.Getpid(), os.Getppid())
	fmt.Printf("uid %v,\tgit %v\n", os.Getuid(), os.Getgid())

	newprocess()
	fmt.Println("end")
	time.Sleep(1 * time.Second)
}

func newcmd() {
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		glog.Println(err)
		return
	}
	//fmt.Println(string(put))
	//cmd.Wait()
}

func newprocess() {

	env := os.Environ()
	fmt.Println("env is: ", env)
	attr := &os.ProcAttr{
		Env: env,
		//Dir:"/home",
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	pid, err := os.StartProcess("./showargs", []string{"ls", "-e"}, attr)
	if err != nil {
		glog.Println(err)
		os.Exit(1)
	}
	pid.Wait()

	fmt.Printf("pid is %v\n", pid)
}
