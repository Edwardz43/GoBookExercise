package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var width = flag.Int("w", 256, "has width(256 or 521)")

func main() {
	flag.Parse()

	var f func(b []byte) []byte

	switch *width {
	case 256:
		f = func(b []byte) []byte {
			h := sha256.Sum256(b)
			return h[:]
		}
	case 512:
		f = func(b []byte) []byte {
			h := sha512.Sum512(b)
			return h[:]
		}
	default:
		log.Fatal("unexpected width specified")
	}
	r := strings.NewReader("test")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", f(b))

}
