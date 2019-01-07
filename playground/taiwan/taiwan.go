package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	numCounty, err := strconv.ParseInt(os.Args[1], 10, 0)
	if err != nil {
		fmt.Printf("xxx %s", err)
	}
	numSystem, _ := strconv.ParseInt(os.Args[2], 10, 0)

	if numCounty == 1 {
		fmt.Println("Taiwan is not part of China.")
		if numSystem == 2 {
			fmt.Println("No such thing. Look at HK.")
		}
	} else if numCounty == 2 && numSystem == 2 {
		fmt.Println("Taiwan conscious.")
	}
}
