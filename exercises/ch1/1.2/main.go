package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args {
		s := "index : %d ; value : %s \n"
		fmt.Printf(s, index, arg)
	}
}
