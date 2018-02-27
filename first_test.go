package scrappy

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestF_Index(t *testing.T) {
	s := New()
	i := 5
	f := s.F.Index(i)
	if f.index != 5 {
		t.Fatal("Unexpected error", f.index, "instead", i)
	}
	doc, err := html.Parse(strings.NewReader(sibling))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.FirstChild.NextSibling.FirstChild.FirstChild.NextSibling
	result := s.F.Index(2).NextSibling(node, Value("xs"))
	if result.FirstChild.Data != "5" {
		t.Fatal("Unexpected error, wrong node")
	}

}

func TestF_Parent(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling
	result := s.F.Parent(node, Tag("body"))
	if result == nil {
		t.Fatal("Unexpected error, node missed")
	}
	result = s.F.Parent(node, Tag("html"), Attr("onload"))
	if result == nil {
		t.Fatal("Unexpected error, node missed")
	}

}

func TestF_Depth(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	// test with tag
	result := s.F.Depth(doc, Tag("p"))
	if result.FirstChild.Data != "a" {
		t.Fatal("Unexpected error wrong node")
	}

	// test with tag, attr and value
	result = s.F.Depth(doc, Tag("a"), Attr("href"), Value("scrappy"))
	if result.FirstChild.Data != "content" {
		t.Fatal("Unexpected error wrong node")
	}

	// test with a valid tag, attr and a non valid value
	result = s.F.Depth(doc, Tag("a"), Attr("href"), Value("scrapp"))
	if result != nil {
		t.Fatal("Unexpected node")
	}

	// test with tag, attr and value
	result = s.F.Depth(doc, Tag("a"), Attr("href"), ContainsValue("scrapp"))
	if result == nil {
		t.Fatal("Unexpected error")
	}

	// test with content
	result = s.F.Depth(doc, Tag("a"), Attr("href"), ContainsValue("scrapp"))
	if result == nil {
		t.Fatal("Unexpected error")
	}
}

func TestF_Breadth(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	// test with tag
	result := s.F.Breadth(doc, Tag("strong"))
	if result.FirstChild.Data != "breadth" {
		t.Fatal("Unexpected error,  wrong node")
	}

	// test with another tag
	result = s.F.Breadth(doc, Tag("p"))
	if !strings.Contains(result.FirstChild.Data, "more") {
		t.Fatal("Unexpected error,  wrong node", result.FirstChild.Data)
	}
}

func TestF_FirstChild(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling
	result := s.F.FirstChild(node, Tag("ul"))
	if result == nil {
		t.Fatal("Unexpected error, node missed")
	}
	node = doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.FirstChild.NextSibling
	result = s.F.FirstChild(node, Tag("section"), Value("xs"))
	if result.Attr[0].Val != "xs" || len(result.Attr) != 1 {
		t.Fatal("Unexpected error, wrong node")
	}
}

func TestF_LastChild(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.FirstChild.NextSibling
	result := s.F.LastChild(node, Tag("section"), Value("xs"))
	if result.Attr[0].Val != "xs" && len(result.Attr) != 2 {
		t.Fatal("Unexpected error, wrong node")
	}
}

func TestF_FirstSibling(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(sibling))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.FirstChild.NextSibling.FirstChild.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling
	result := s.F.FirstSibling(node, Value("md"))
	if result.FirstChild.Data != "2" {
		t.Fatal("Unexpected error, wrong node")
	}
	result = s.F.FirstSibling(node, Value("lg"))
	if result != nil {
		t.Fatal("Unexpected error")
	}
	result = s.F.FirstSibling(node, Value("xs"))
	if result.FirstChild.Data != "1" {
		t.Fatal("Unexpected error, wrong node")
	}
}

func TestF_LastSibling(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(sibling))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.FirstChild.NextSibling.FirstChild.FirstChild.NextSibling
	result := s.F.LastSibling(node, Value("xs"))
	if result.FirstChild.Data != "7" {
		t.Fatal("Unexpected error, wrong node")
	}
}

func TestF_NextSibling(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(sibling))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.FirstChild.NextSibling.FirstChild.FirstChild.NextSibling
	result := s.F.NextSibling(node, Value("xs"))
	if result.FirstChild.Data != "3" {
		t.Fatal("Unexpected error, wrong node")
	}
}

func TestF_PrevSibling(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(sibling))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	node := doc.FirstChild.FirstChild.NextSibling.FirstChild.FirstChild.NextSibling
	result := s.F.NextSibling(node, Value("xs"))
	if result.FirstChild.Data != "3" {
		t.Fatal("Unexpected error, wrong node")
	}
	result = s.F.PrevSibling(result, Value("xs"))
	if result.FirstChild.Data != "1" {
		t.Fatal("Unexpected error, wrong node")
	}
}
