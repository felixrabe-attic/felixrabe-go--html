package html

import (
	"bytes"
	"io"
	"net/http"
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

func ParseUrl(url string) (*Node, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	doc, err := Parse(r.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
