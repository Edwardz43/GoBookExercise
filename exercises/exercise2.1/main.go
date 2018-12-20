package main

import (
	"GoBook/book/ch2/tempconv"
	"fmt"
)

func main() {
	c := tempconv.AbsoluteZeroC
	k := tempconv.AbsoluteZeroK

	KToC := tempconv.KToC

	fmt.Println(c == KToC(k)) //true
	fmt.Printf("%g°C is equal %g°K\n", c, k)
}
