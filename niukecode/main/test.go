package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func testcompare() {

	fmt.Printf("1 input:\n")
	readb := bufio.NewReader(os.Stdin)
	s, isnext, err := readb.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("next is %v ,s is %v\n", isnext, s)
	fmt.Printf("2 input:\n")
	b, isnext, err := readb.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	if !bytes.Equal(s, b) {
		for y := 0; y < len(s); y++ {
			if s[y] != b[y] {
				fmt.Printf(" not equal is %d,%d,%d", y, s[y], b[y])
			}
		}
	} else {
		fmt.Printf("equal!\n")
	}
}
