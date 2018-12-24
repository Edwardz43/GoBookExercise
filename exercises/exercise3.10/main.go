package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}

}

func comma(s string) string {

	var b bytes.Buffer

	pre := len(s) % 3

	if pre == 0 {
		pre = 3
	}

	// filter the prefix 0
	var n int

	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			break
		}
		n++
	}

	b.WriteString(s[n:pre])

	for i := pre; i < len(s); i += 3 {
		if b.Len() != 0 {
			b.WriteByte(',')
		}

		b.WriteString(s[i : i+3])
	}

	// convert 00000... to 0
	if b.Len() == 0 {
		return "0"
	}

	return b.String()
}
