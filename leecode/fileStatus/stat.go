package fileStatus

import (
	"fmt"
	"os"
	"os/exec"
)

func GetFileStat(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	//os.StartProcess()
	fmt.Printf("%d\n", info.Mode().Perm()&os.ModeSetuid)
	fmt.Printf("name is %v,size %v,mode %v\n", info.Name(), info.Size(), info.Mode().String())
	return nil
}

func Getid() error {
	fmt.Printf("uid is %v,euid is %v\n", os.Getuid(), os.Geteuid())
	d, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
	niu := exec.Command("E:\\GoFile\\bin\\niuke.exe")
	niu.Stderr = os.Stderr
	niu.Stdout = os.Stdout
	niu.Stdin = os.Stdin
	fmt.Println(niu.Run())

	return nil
}
