package main

import (
	"GoBook/book/ch2/popcount"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseInt(arg, 10, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		fmt.Println(popcount.PopCount(uint64(x)))
		fmt.Println(popcount.V2(uint64(x)))
	}

}
