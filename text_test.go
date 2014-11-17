package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html/atom"
)

func TestGetAllText(t *testing.T) {
	doc, err := ParseString("<p>  This is <i>some</i> text for <strong>extra <em>emphasis!</em></strong></p>")
	assert.NoError(t, err)
	p := doc.FindByAtom(atom.Html, atom.Body, atom.P)[0]
	assert.Equal(t, "  This is some text for extra emphasis!", p.GetAllText())
}
