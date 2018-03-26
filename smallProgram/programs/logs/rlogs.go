package main

import (
	"errors"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.Println("ok")
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)

	l := []int{0, 1, 2}
	n, err := pop(&l)
	for i := 0; i < 5 && err == nil; i++ {
		logrus.Printf("i is %v,n is %v ", i, n)
		logrus.Printf("l is %v\n", l)
		n, err = pop(&l)
	}

	if err != nil {
		logrus.Println(err)
	}

	log.Println("ok")

}

func pop(l *[]int) (int, error) {
	if len(*l) == 0 {
		return 0, errors.New("nil silce")
	}
	r := (*l)[0]
	*l = (*l)[1:]
	return r, nil

}
