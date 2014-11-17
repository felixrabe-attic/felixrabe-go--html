package html

import (
	"golang.org/x/net/html/atom"
)

// FindByAtom finds element nodes by their atom across multiple levels.
func (node *Node) FindByAtom(atoms ...atom.Atom) (nodes []*Node) {
	nodes = make([]*Node, 1)
	nodes[0] = node
	for _, a := range atoms {
		children := make([]*Node, 0)
		for _, n := range nodes {
			for _, c := range n.ElementChildren() {
				if c.DataAtom == a {
					children = append(children, c)
				}
			}
		}
		nodes = children
	}
	return
}

// NextByAtom finds the next sibling element by its atom.
func (node *Node) NextByAtom(atom atom.Atom) *Node {
	for node = node.NextElement(); node != nil; node = node.NextElement() {
		if node.DataAtom == atom {
			return node
		}
	}
	return nil
}

// PrevByAtom finds the previous sibling element by its atom.
func (node *Node) PrevByAtom(atom atom.Atom) *Node {
	for node = node.PrevElement(); node != nil; node = node.PrevElement() {
		if node.DataAtom == atom {
			return node
		}
	}
	return nil
}
