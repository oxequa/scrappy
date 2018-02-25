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
	result := s.All.Child(node, Tag("div"))
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
	result := s.All.Depth(doc, ContainsText("more"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'more' instead", len(result))
	}

	result = s.All.Depth(doc, Text("content"))
	if len(result) != 3 {
		t.Fatal("Expected 3 nodes 'content' instead", len(result))
	}

	result = s.All.Depth(doc, Text("href"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'content' instead", len(result))
	}

	// test with tag
	result = s.All.Depth(doc, Tag("a"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'a' instead", len(result))
	}
	result = s.All.Depth(doc, Tag("div"))
	if len(result) != 3 {
		t.Fatal("Expected 3 nodes 'div' instead", len(result))
	}
	result = s.All.Depth(doc, Tag("href"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'href' instead", len(result))
	}

	// test with attr
	result = s.All.Depth(doc, Attr("href"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'href' instead", len(result))
	}
	result = s.All.Depth(doc, Attr("onload"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'onload' instead", len(result))
	}
	result = s.All.Depth(doc, Attr("div"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'div' instead", len(result))
	}

	// test with attr value
	result = s.All.Depth(doc, Value("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'myFunc' instead", len(result))
	}
	result = s.All.Depth(doc, Value("scrappy"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'scrappy' instead", len(result))
	}

	// mixing
	result = s.All.Depth(doc, Attr("onload"), Value("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}
	result = s.All.Depth(doc, Tag("a"), Attr("href"), Value("scrappy"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}

	// mixing with custom filter
	custom := func(node *html.Node) bool {
		result := s.All.Depth(node, Text("content"))
		return len(result) > 0
	}
	s.All.Depth(doc, Tag("a"), Attr("href"), Value("scrappy"), custom)
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}

	// nested filters enabled
	s.nested = true
	result = s.All.Depth(doc, Tag("body"), Tag("div"), Tag("section"), Attr("href"), Text("more"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}
	result = s.All.Depth(doc, Tag("body"), Tag("div"), Tag("section"), Attr("p"), Text("more"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes instead", len(result))
	}
	result = s.All.Depth(doc, Tag("html"), Attr("onload"), Value("myFunc()"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes instead", len(result))
	}

	// nested filters disabled
	s.nested = false
	result = s.All.Depth(doc, Tag("html"), Attr("onload"), Value("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}
	result = s.All.Depth(doc, Tag("body"), Tag("div"), Tag("section"), Attr("href"), Text("more"))
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
	result := s.All.Parent(node)
	if len(result) != 3 {
		t.Fatal("Unexpected error, there are 3 parent instead", len(result))
	}
	node = doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.LastChild
	result = s.All.Parent(node)
	if len(result) != 2 {
		t.Fatal("Unexpected error, there are 2 parent instead", len(result))
	}
	result = s.All.Parent(node, Tag("body"))
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
	result := s.All.Breadth(doc, ContainsText("more"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'more' instead", len(result))
	}
	result = s.All.Breadth(doc, Text("content"))
	if len(result) != 3 {
		t.Fatal("Expected 3 nodes 'content' instead", len(result))
	}
	result = s.All.Breadth(doc, Text("href"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'content' instead", len(result))
	}

	// test with tag
	result = s.All.Breadth(doc, Tag("a"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'a' instead", len(result))
	}
	result = s.All.Breadth(doc, Tag("div"))
	if len(result) != 3 {
		t.Fatal("Expected 3 nodes 'div' instead", len(result))
	}
	result = s.All.Breadth(doc, Tag("href"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'href' instead", len(result))
	}

	// test with attr
	result = s.All.Breadth(doc, Attr("href"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'href' instead", len(result))
	}
	result = s.All.Breadth(doc, Attr("onload"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'onload' instead", len(result))
	}
	result = s.All.Breadth(doc, Attr("div"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'div' instead", len(result))
	}

	// test with attr value
	result = s.All.Breadth(doc, Value("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'myFunc' instead", len(result))
	}
	result = s.All.Breadth(doc, Value("scrappy"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'scrappy' instead", len(result))
	}
	// mixing
	result = s.All.Breadth(doc, Attr("onload"), Value("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}
	result = s.All.Breadth(doc, Tag("a"), Attr("href"), Value("scrappy"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead", len(result))
	}

	// mixing with custom filter
	custom := func(node *html.Node) bool {
		result := s.All.Breadth(node, Text("content"))
		return len(result) > 0
	}
	s.All.Breadth(doc, Tag("a"), Attr("href"), Value("scrappy"), custom)
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
	result := s.All.NextSibling(node, Tag("section"))
	if len(result) != 1 {
		t.Fatal("Unexpected error, expected only one section node instead", len(result))
	}
	node = doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.FirstChild.NextSibling
	result = s.All.NextSibling(node, Tag("div"))
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
	result := s.All.PrevSibling(node, Tag("section"))
	if len(result) != 1 {
		t.Fatal("Unexpected error, expected only one section node instead", len(result))
	}
	node = doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.LastChild.PrevSibling
	result = s.All.PrevSibling(node, Tag("div"))
	if len(result) != 2 {
		t.Fatal("Unexpected error, expected there are 2 div node instead", len(result))
	}
}
