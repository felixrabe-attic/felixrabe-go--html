// +build ignore

package html

func HtmlChildren(node *html.Node) (children []*html.Node) {
	children = make([]*html.Node, 0)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		children = append(children, child)
	}
	return
}

func HtmlElementChildren(node *html.Node) (children []*html.Node) {
	children = make([]*html.Node, 0)
	for _, c := range HtmlChildren(node) {
		if c.Type != html.ElementNode {
			continue
		}
		children = append(children, c)
	}
	return
}

func HtmlFindByAtom(node *html.Node, atoms ...atom.Atom) (nodes []*html.Node) {
	nodes = make([]*html.Node, 1)
	nodes[0] = node
	for _, a := range atoms {
		children := make([]*html.Node, 0)
		for _, n := range nodes {
			for _, c := range HtmlElementChildren(n) {
				if c.DataAtom != a {
					continue
				}
				children = append(children, c)
			}
		}
		nodes = children
	}
	return
}

func HtmlNextByAtom(node *html.Node, atom atom.Atom) *html.Node {
	node = node.NextSibling

	for node != nil {
		if node.Type == html.ElementNode {
			if node.DataAtom == atom {
				return node
			}
		}
	}
	return nil
}

func HtmlText(node *html.Node) string {
	strs := make([]string, 0)
	for _, c := range HtmlChildren(node) {
		switch c.Type {
		case html.TextNode:
			strs = append(strs, c.Data)
		case html.ElementNode:
			strs = append(strs, HtmlText(c))
		}
	}
	return strings.Join(strs, "")
}

func HtmlNormalText(node *html.Node) string {
	return strings.Join(strings.Fields(HtmlText(node)), " ")
}

func HtmlFilterNormalText(nodes []*html.Node, fn func(string) bool) (nn []*html.Node) {
	nn = make([]*html.Node, 0)
	for _, node := range nodes {
		if fn(HtmlNormalText(node)) {
			nn = append(nn, node)
		}
	}
	return
}

func HtmlStringify(node *html.Node) (text string) {
	b := &bytes.Buffer{}
	if err := html.Render(b, node); err != nil {
		misc.FatalIf(err)
	}
	return b.String()
}
