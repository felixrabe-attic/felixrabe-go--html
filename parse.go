package html

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func Parse(r io.Reader) (*Node, error) {
	doc, err := html.Parse(r)
	return (*Node)(doc), err
}

func ParseBytes(b []byte) (*Node, error) {
	return Parse(bytes.NewBuffer(b))
}

func ParseString(s string) (*Node, error) {
	return Parse(strings.NewReader(s))
}
