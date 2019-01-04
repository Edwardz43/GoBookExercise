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

	links := make(map[string][]string)

	for k, links := range visit(links, doc) {
		fmt.Printf("Element Type : %v\n", k)
		for _, link := range links {
			fmt.Println(link)
		}
		fmt.Println("\n --------------------------------")
	}
}

func visit(links map[string][]string, n *html.Node) map[string][]string {
	if n.Type == html.ElementNode {
		for _, e := range n.Attr {
			links[n.Data] = append(links[n.Data], e.Val)
		}

	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}
