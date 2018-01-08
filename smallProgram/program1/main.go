package main

import (
	"fmt"
	"os"
)



func main() {
	fmt.Printf("argnum %d,args %v\n ",len(os.Args),os.Args)
}

