package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var (
		err  error
		path = "/home/job.log"
		i    = 1
	)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_SYNC, 0777)
	if err != nil {
		log.Println(err)
		return
	}

	for ; i <= 100; i++ {
		if _, err = fmt.Fprintf(f, "%d\n", i); err != nil {
			log.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
	}

	if err = f.Sync(); err != nil {
		log.Println(err)
	}
	f.Close()
}
