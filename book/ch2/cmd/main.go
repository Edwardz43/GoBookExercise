package main

import (
	"fmt"
	"gobook/book/ch2/tempconv"
)

func main() {
	var f tempconv.Fahrenheit

	f = 212

	c := tempconv.FToC(f)

	fmt.Println(c)

	b := tempconv.BoilingC

	fmt.Println(b)

	fmt.Println(tempconv.CToF(b))
}
