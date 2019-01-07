package outline2

import (
	"fmt"

	"golang.org/x/net/html"
)

func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

// func main() {
// 	doc, err := html.Parse(os.Stdin)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
// 		os.Exit(1)
// 	}

// 	ForEachNode(doc, startElement, endElement)
// }
