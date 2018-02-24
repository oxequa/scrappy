package scrappy

import (
	"testing"
	"golang.org/x/net/html"
	"github.com/go-siris/siris/core/router/node"
)

func TestFirst_Index(t *testing.T) {
	s := New()
	i := 5
	f := s.First.Index(i)
	if f.index != 5 {
		t.Fatal("Unexpected error", f.index, "instead", i)
	}
}

func TestFirst_Child(t *testing.T) {
	var result []*html.Node
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		if a.Validate(node, filters...) {
			result = append(result,node)
		}
	}
	return result
}
