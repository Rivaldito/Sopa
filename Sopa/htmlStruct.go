package sopa

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type HTMLNode struct {
	Pointer   *html.Node
	NodeValue string
	Error     error
}

func (sopa Sopa) HTMLParse() HTMLNode {

	htmlParse, err := html.Parse(strings.NewReader(*sopa.HTMLResponse))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return HTMLNode{
			Error: newError(ErrUnableToParse, "unable to parse the HTML"),
		}
	}

	for htmlParse.Type != html.ElementNode {
		switch htmlParse.Type {
		case html.DocumentNode:
			htmlParse = htmlParse.FirstChild
		case html.DoctypeNode:
			htmlParse = htmlParse.NextSibling
		case html.CommentNode:
			htmlParse = htmlParse.NextSibling
		}
	}

	node := HTMLNode{
		Pointer:   htmlParse,
		NodeValue: htmlParse.Data,
		Error:     nil,
	}

	return node
}
