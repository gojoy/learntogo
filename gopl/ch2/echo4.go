package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit newline")
var sep = flag.String("s", "", "separator")

func main() {
	flag.Parse()
	fmt.Println("args is ", flag.Args())
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
