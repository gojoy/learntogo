package dps

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const MAX = 32

func Testio() {

	var buf bytes.Buffer
	tmp := make([]byte, 4)

	fmt.Printf("len is %d\n", len(tmp))
	n, err := os.Stdin.Read(tmp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("read  %v byte,%c\n", n, tmp[0])
	buf.Write(tmp)
	fmt.Printf("buf is %v\n", buf.String())

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("input:\n")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		println(err)
	} else {
		fmt.Printf("%v\n", input)
	}
}

func DirectIo() {

	buf := make([]byte, MAX)

	fmt.Printf("begin\n")

	_, err := os.Stdin.Read(buf)

	for err == nil {
		if _, err := os.Stdout.Write(buf); err != nil {
			panic("write error " + err.Error())
		}
		_, err = os.Stdin.Read(buf)
	}

	fmt.Println("read error" + err.Error())

}

func Buffio() {

	buf := make([]byte, MAX)
	inputReader := bufio.NewReader(os.Stdin)
	outputWriter := bufio.NewWriter(os.Stdout)

	for {
		_, rerr := inputReader.Read(buf)
		if rerr != nil {
			fmt.Printf("read err: %v\n ", rerr.Error())
			break
		}
		fmt.Printf("buf is %v\n", buf)
		if _, werr := outputWriter.Write(buf); werr != nil {
			panic("write error:" + werr.Error())
		}
	}
	if err := outputWriter.Flush(); err != nil {
		panic(err)
	}

}
