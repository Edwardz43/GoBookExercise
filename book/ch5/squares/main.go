package main

import "fmt"

func spuares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {

	f := spuares()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
