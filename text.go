package html

import (
	"strings"

	"golang.org/x/net/html"
)

// GetAllText gets all text without any tags.
func (node *Node) GetAllText() string {
	text := make([]string, 0)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		switch child.Type {
		case html.ElementNode:
			text = append(text, (*Node)(child).GetAllText())
		case html.TextNode:
			text = append(text, child.Data)
		}
	}
	return strings.Join(text, "")
}
