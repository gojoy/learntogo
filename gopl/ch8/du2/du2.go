package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var (
		starttime = time.Now()
		nfiles    int64
		nbytes    int64
		fsize     = make(chan int64)
		startdir  = flag.String("dir", "E:\\GoFile\\src\\learntogo", "set dir")
		verbose   = flag.Bool("v", false, "show verbose message")
		tick      = make(<-chan time.Time)
	)

	flag.Parse()
	fmt.Println("startdir is ", *startdir)
	if ss, err := os.Stat(*startdir); err != nil || !ss.IsDir() {
		log.Fatalln(*startdir, " is not a dir")
	}
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	go func() {
		walkDire(*startdir, fsize)
		close(fsize)
	}()

	for {
		select {
		case s, ok := <-fsize:
			if !ok {
				goto end
				panic("recive done\n")
			}
			nfiles++
			nbytes += s
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
end:
	printDiskUsage(nfiles, nbytes)
	fmt.Printf("cost %v\n", time.Since(starttime))
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
