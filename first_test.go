package scrappy

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
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
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling
	result := s.First.Child(node, Tag("ul"))
	if result == nil {
		t.Fatal("Unexpected error, node missed")
	}
	node = doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.FirstChild.NextSibling
	result = s.First.Child(node, Tag("section"), ContainsValue("xs"), Value("col xs"))
	if result == nil {
		t.Fatal("Unexpected error, node missed")
	}
}

func TestFirst_Parent(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling
	result := s.First.Parent(node, Tag("body"))
	if result == nil {
		t.Fatal("Unexpected error, node missed")
	}
	result = s.First.Parent(node, Tag("html"), Attr("onload"))
	if result == nil {
		t.Fatal("Unexpected error, node missed")
	}

}

func TestFirst_Depth(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	result := s.First.Depth(doc, Tag("p"))
	if result.FirstChild.Data != "a" {
		t.Fatal("Unexpected error wrong node")
	}
	result = s.First.Depth(doc, Tag("a"), Attr("href"), Value("scrappy"))
	if result.FirstChild.Data != "content" {
		t.Fatal("Unexpected error wrong node")
	}
}

func TestFirst_Breadth(t *testing.T) {

}
