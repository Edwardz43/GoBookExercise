package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("hello world")

	fmt.Println(b)

	b = revUTF8(b)

	fmt.Println(b)

	s := string(b)

	fmt.Println(s)
}

func rev(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

// ex4.7
// Reverse all the runes, and then the entire slice. The runes' bytes end up in
// the right order.
func revUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size])
		i += size
	}
	rev(b)
	return b
}
