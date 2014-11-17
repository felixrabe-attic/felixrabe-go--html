package html

import (
	"bytes"
	"fmt"

	"golang.org/x/net/html"
)

func (node *Node) String() string {
	b := &bytes.Buffer{}
	if err := html.Render(b, (*html.Node)(node)); err != nil {
		return "<<ERROR>>"
	}
	return b.String()
}

func (node *Node) RawString() string {
	return fmt.Sprint((*html.Node)(node))
}
