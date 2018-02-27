package scrappy

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestAll_Child(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling
	result := s.A.Child(node, Tag("div"))
	if len(result) != 3 {
		t.Fatal("Unexpected error, expected three div node instead", len(result))
	}
}

func TestAll_Depth(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	// test with content
	result := s.A.Depth(doc, ContainsText("more"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'more' instead", len(result))
	}

	result = s.A.Depth(doc, Text("content"))
	if len(result) != 3 {
		t.Fatal("Expected 3 nodes 'content' instead", len(result))
	}

	result = s.A.Depth(doc, Text("href"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'content' instead", len(result))
	}

	// test with tag
	result = s.A.Depth(doc, Tag("a"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'a' instead", len(result))
	}
	result = s.A.Depth(doc, Tag("div"))
	if len(result) != 3 {
		t.Fatal("Expected 3 nodes 'div' instead", len(result))
	}
	result = s.A.Depth(doc, Tag("href"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'href' instead", len(result))
	}

	// test with attr
	result = s.A.Depth(doc, Attr("href"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'href' instead", len(result))
	}
	result = s.A.Depth(doc, Attr("onload"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'onload' instead", len(result))
	}
	result = s.A.Depth(doc, Attr("div"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'div' instead", len(result))
	}

	// test with attr value
	result = s.A.Depth(doc, Value("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'myFunc' instead", len(result))
	}
	result = s.A.Depth(doc, Value("scrappy"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'scrappy' instead", len(result))
	}

	// mixing
	result = s.A.Depth(doc, Attr("onload"), Value("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}
	result = s.A.Depth(doc, Tag("a"), Attr("href"), Value("scrappy"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}

	// mixing with custom filter
	custom := func(node *html.Node) bool {
		result := s.A.Depth(node, Text("content"))
		return len(result) > 0
	}
	s.A.Depth(doc, Tag("a"), Attr("href"), Value("scrappy"), custom)
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}

	// nested filters enabled
	s.nested = true
	result = s.A.Depth(doc, Tag("body"), Tag("div"), Tag("section"), Attr("href"), Text("more"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}
	result = s.A.Depth(doc, Tag("body"), Tag("div"), Tag("section"), Attr("p"), Text("more"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes instead", len(result))
	}
	result = s.A.Depth(doc, Tag("html"), Attr("onload"), Value("myFunc()"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes instead", len(result))
	}

	// nested filters disabled
	s.nested = false
	result = s.A.Depth(doc, Tag("html"), Attr("onload"), Value("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}
	result = s.A.Depth(doc, Tag("body"), Tag("div"), Tag("section"), Attr("href"), Text("more"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes instead", len(result))
	}
}

func TestAll_Parent(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling
	result := s.A.Parent(node, Tag("body"))
	if len(result) != 1 {
		t.Fatal("Unexpected error, there is 1 parent body instead", len(result))
	}
}

func TestAll_Breadth(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	// test with content
	result := s.A.Breadth(doc, ContainsText("more"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'more' instead", len(result))
	}
	result = s.A.Breadth(doc, Text("content"))
	if len(result) != 3 {
		t.Fatal("Expected 3 nodes 'content' instead", len(result))
	}
	result = s.A.Breadth(doc, Text("href"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'content' instead", len(result))
	}

	// test with tag
	result = s.A.Breadth(doc, Tag("a"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'a' instead", len(result))
	}
	result = s.A.Breadth(doc, Tag("div"))
	if len(result) != 3 {
		t.Fatal("Expected 3 nodes 'div' instead", len(result))
	}
	result = s.A.Breadth(doc, Tag("href"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'href' instead", len(result))
	}

	// test with attr
	result = s.A.Breadth(doc, Attr("href"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'href' instead", len(result))
	}
	result = s.A.Breadth(doc, Attr("onload"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'onload' instead", len(result))
	}
	result = s.A.Breadth(doc, Attr("div"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'div' instead", len(result))
	}

	// test with attr value
	result = s.A.Breadth(doc, Value("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'myFunc' instead", len(result))
	}
	result = s.A.Breadth(doc, Value("scrappy"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'scrappy' instead", len(result))
	}
	// mixing
	result = s.A.Breadth(doc, Attr("onload"), Value("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}
	result = s.A.Breadth(doc, Tag("a"), Attr("href"), Value("scrappy"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}

	// mixing with custom filter
	custom := func(node *html.Node) bool {
		result := s.A.Breadth(node, Text("content"))
		return len(result) > 0
	}
	s.A.Breadth(doc, Tag("a"), Attr("href"), Value("scrappy"), custom)
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}
}

func TestAll_NextSibling(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling
	result := s.A.NextSibling(node, Tag("section"))
	if len(result) != 3 {
		t.Fatal("Unexpected error, expected three section node instead", len(result))
	}
	node = doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.FirstChild.NextSibling
	result = s.A.NextSibling(node, Tag("div"))
	if len(result) != 2 {
		t.Fatal("Unexpected error, expected 2 section node instead", len(result))
	}
}

func TestAll_PrevSibling(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.FirstChild.NextSibling.LastChild.PrevSibling
	result := s.A.PrevSibling(node, Tag("section"))
	if len(result) != 3 {
		t.Fatal("Unexpected error, expected only three section node instead", len(result))
	}
	node = doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.LastChild.PrevSibling
	result = s.A.PrevSibling(node, Tag("div"))
	if len(result) != 2 {
		t.Fatal("Unexpected error, expected there are 2 div node instead", len(result))
	}
}
