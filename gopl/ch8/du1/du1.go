package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var (
		sizesum  int64
		filenum  int64
		fsize    = make(chan int64)
		startdir = flag.String("dir", "E:\\GoFile\\src\\learntogo", "set dir")
	)

	flag.Parse()
	fmt.Println("startdir is ", *startdir)
	ss, err := os.Stat(*startdir)
	if err != nil || !ss.IsDir() {
		log.Fatalln(*startdir, " is not a dir")
	}
	go func() {
		walkDire(*startdir, fsize)
		close(fsize)
	}()
	for s := range fsize {
		filenum++
		sizesum += s
	}
	printDiskUsage(filenum, sizesum)
}

func walkDire(dir string, filesize chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDire(subdir, filesize)
		} else {
			filesize <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("readDir error: ", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files\t%.2f MB\n", nfiles, float64(nbytes)/1e6)
}
