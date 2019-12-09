package main

import (
	"fmt"
	"github.com/Edwardz43/GoBookExercise/book/ch2/tempconv"
)

func main() {
	c := tempconv.AbsoluteZeroC
	k := tempconv.AbsoluteZeroK

	KToC := tempconv.KToC

	fmt.Println(c == KToC(k)) //true
	fmt.Printf("%g°C is equal %g°K\n", c, k)
}
