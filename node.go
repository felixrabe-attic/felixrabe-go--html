package html

import (
	"golang.org/x/net/html"
)

// Children returns all child nodes.
func (node *Node) Children() (children []*Node) {
	children = make([]*Node, 0)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		children = append(children, (*Node)(child))
	}
	return
}

// ElementChildren returns all child nodes that are elements.
func (node *Node) ElementChildren() (children []*Node) {
	children = make([]*Node, 0)
	for _, c := range node.Children() {
		if c.Type == html.ElementNode {
			children = append(children, c)
		}
	}
	return
}

func (node *Node) Next() *Node {
	return (*Node)(node.NextSibling)
}

func (node *Node) NextElement() *Node {
	for node = node.Next(); node != nil; node = node.Next() {
		if node.Type == html.ElementNode {
			return node
		}
	}
	return nil
}

func (node *Node) Prev() *Node {
	return (*Node)(node.PrevSibling)
}

func (node *Node) PrevElement() *Node {
	for node = node.Prev(); node != nil; node = node.Prev() {
		if node.Type == html.ElementNode {
			return node
		}
	}
	return nil
}

func (node *Node) RemoveChild(child *Node) {
	(*html.Node)(node).RemoveChild((*html.Node)(child))
}

func (node *Node) AppendChild(child *Node) {
	(*html.Node)(node).AppendChild((*html.Node)(child))
}
