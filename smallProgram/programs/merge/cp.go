package main

import (
	"log"
	"os/exec"
)

func main() {
	var (
		err error
		src = "/var/dir1/"
		dst = "/var/dir2/"
	)
	s := append([]byte(src), 42)
	if err = cpdir(string(s), dst); err != nil {
		log.Println(err)
	}
	return
}

func cpdir(src, dst string) error {
	args := []string{"-rf"}
	args = append(args, src, dst)
	cmd := exec.Command("cp", args...)

	buf, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("err is %v,out is %v\n", err, string(buf))
		return err
	}
	return nil
}
