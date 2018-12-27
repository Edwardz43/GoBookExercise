package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("a"))
	c2 := sha256.Sum256([]byte("A"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	fmt.Println(checkDiff(c1, c2))
}

func checkDiff(c1 [32]byte, c2 [32]byte) int {
	count := 0
	for i, v := range c1 {
		count += popCount(v ^ c2[i])
	}
	return count
}

func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}
