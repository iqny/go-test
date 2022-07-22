package main

import (
	"fmt"
	"golang.org/x/net/html"
)

func main() {
	fmt.Printf("'%*s' %s",4,"","abc")
}

func forEachNode(n *html.Node, per, post func(n *html.Node)) {
	if per != nil {
		per(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, per, post)
	}
	if post != nil {
		post(n)
	}
}
