package html

import (
	"golang.org/x/net/html"
)

func (node *Node) RemoveAttr(key string) {
	newAttr := make([]html.Attribute, 0)
	for _, a := range node.Attr {
		if !(a.Namespace == "" && a.Key == key) {
			newAttr = append(newAttr, a)
		}
	}
	node.Attr = newAttr
}
