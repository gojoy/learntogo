package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type dep [2]uint64

func main() {

	var (
		flag         bool = true
		err, errconv error
	)
	depall := make([]dep, 0)

	inputReader := bufio.NewReader(os.Stdin)
	str, err := inputReader.ReadString('\n')
	for err == nil {
		str = str[:len(str)-1]
		if str[len(str)-1] != ',' {
			flag = false
		} else {
			str = str[:len(str)-1]
		}
		strs := strings.Split(str, ",")
		first := strings.TrimSpace(strs[0][1:])
		last := strings.TrimSpace(strs[1][:len(strs[1])-1])
		depall, errconv = inputHandle(first, last, depall)
		if errconv != nil {
			fmt.Println(errconv)
			return
		}
		if !flag {
			break
		}
		str, err = inputReader.ReadString('\n')
	}
	if err != nil && err != io.EOF {
		fmt.Printf("input error: %v\n", err)
	}
	fmt.Println(depall)
}

func inputHandle(f, l string, a []dep) ([]dep, error) {
	var (
		err    error
		fn, ln uint64
	)
	fn, err = strconv.ParseUint(f, 0, 0)
	if err != nil {
		return a, err
	}
	ln, err = strconv.ParseUint(l, 0, 0)
	if err != nil {
		return a, err
	}
	a = append(a, dep{fn, ln})
	return a, nil
}
