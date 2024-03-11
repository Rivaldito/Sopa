package sopa

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func (node *HTMLNode) HTMLPrint(depth int) {

	for node.Pointer != nil {

		indent := strings.Repeat("  ", depth)

		switch node.Pointer.Type {
		case html.ElementNode:
			fmt.Printf("%s<%s>\n", indent, node.Pointer.Data)
		case html.TextNode:
			text := strings.TrimSpace(node.Pointer.Data)
			if text != "" {
				fmt.Printf("%s%s\n", indent, text)
			}
		case html.CommentNode:
			fmt.Printf("%s<!-- %s -->\n", indent, node.Pointer.Data)
		}

		node.Pointer = node.Pointer.FirstChild
		node.Pointer = node.Pointer.NextSibling

		depth = depth + 1

	}

}

func (sopa Sopa) TraverseNode(n *html.Node, depth int) {

	indent := strings.Repeat("  ", depth)

	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%s<%s>\n", indent, n.Data)
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Printf("%s%s\n", indent, text)
		}
	case html.CommentNode:
		fmt.Printf("%s<!-- %s -->\n", indent, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sopa.TraverseNode(c, depth+1)
	}
}
