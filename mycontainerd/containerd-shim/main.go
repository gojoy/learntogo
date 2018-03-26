// +build linux

package main

import (
	"flag"
	"fmt"
	"github.com/docker/containerd/osutils"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func newLogger(name string) *log.Logger {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(filepath.Join(cwd, name), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(f, "", log.Ltime|log.Lshortfile)
}

var glog *log.Logger

func init() {

	glog = newLogger("ggshim.log")
}

// shim containerID bundledir runcbinary run/create
func main() {
	flag.Parse()
	if err := start(); err != nil {
		if err == errRuntime {
			glog.Printf("runtime error: %v\n", err)
			return
		}
		glog.Printf("start err: %v\n", err)
		os.Exit(1)
	}
}

func start() error {
	signals := make(chan os.Signal, 2048)
	signal.Notify(signals)

	if err := osutils.SetSubreaper(1); err != nil {
		return err
	}
	p, err := newProcess(flag.Arg(0), flag.Arg(1), flag.Arg(2), flag.Arg(3))
	if err != nil {
		return err
	}
	defer func() {
		if err := p.Close(); err != nil {
			glog.Printf("close err:%v\n", err)
		}
	}()

	if err := p.create(); err != nil {
		glog.Println(err)
		p.delete()
		return err
	}
	var exitShim bool

	for {

		s := <-signals
		switch s {
		case syscall.SIGCHLD:
			glog.Println("recv signal SIGCHLD")
			exits, _ := osutils.Reap(false)
			for _, e := range exits {
				glog.Println("pid", e.Pid, " status", e.Status)
				if e.Pid == p.pid() {
					glog.Println(e.Pid, e.Status)
					exitShim = true
					writeInt("exitStatus", e.Status)
				}
			}
		case syscall.SIGINT:
			glog.Println("recv signal SIGINT")
			os.Exit(1)
		}

		if exitShim {
			glog.Println("exitShim")
			osutils.Reap(true)
			p.Wait()
		}
	}
	return nil
}

func writeInt(path string, i int) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "%d", i)
	return err
}
