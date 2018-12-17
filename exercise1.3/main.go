package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Environ() {
		s += sep + arg
		sep = "~~~~"
	}
	fmt.Println(s)

	fmt.Println("-------------------------------------------------")

	fmt.Println(strings.Join(os.Environ(), "~~~~"))
}
