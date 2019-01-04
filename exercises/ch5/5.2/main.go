package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	ecounter := make(map[string]int)

	for k, v := range visit(ecounter, doc) {
		fmt.Printf("%-10s : %3d\n", k, v)
	}
}

func visit(counter map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		for i := 0; i < len(n.Attr); i++ {
			counter[n.Data]++

		}
	}
	if n.FirstChild != nil {
		counter = visit(counter, n.FirstChild)
	}

	if n.NextSibling != nil {
		counter = visit(counter, n.NextSibling)
	}

	return counter
}
