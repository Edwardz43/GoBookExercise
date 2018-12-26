package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Println(checkDiff(c1, c2))
}

func checkDiff(c1 [32]byte, c2 [32]byte) int {
	count := 0
	for i, v := range c1 {
		if v != c2[i] {
			count++
		}
	}
	return count
}
