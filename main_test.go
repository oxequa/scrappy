package scrappy

import (
	"testing"
	"strings"
	"golang.org/x/net/html"
)

const(
	url = "https://github.com/tockins"
	document = `<!DOCTYPE html>
	<html onload="myFunc()">
	<head>
		<title></title>
	</head>
	<body>
		body text
		<div>
			<p>
				<a href="">more</a>
				<a href="">content</a>
				<div>i am also a more node</a>
			</p>
		</div>
		<p>more <a href="scrappy">content</a></p>
		<p><a href="">content</a> more onload</p>
	</body>
	</html>`
)

func TestGet(t *testing.T) {
	s, err := Get(url)
	if err != nil{
		t.Fatal("Unexpected error", err)
	}
	if s.Root == nil{
		t.Fatal("Unexpected error root node doesn't exist")
	}
}

func TestScrappy_Depth(t *testing.T) {
	s := Scrappy{}
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error",err)
	}
	// test with content
	result := s.Depth(doc,Text("more"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'more' instead",len(result))
	}
	s = Scrappy{}
	result = s.Depth(doc,Text("content"))
	if len(result) != 3 {
		t.Fatal("Expected 3 nodes 'content' instead",len(result))
	}
	s = Scrappy{}
	result = s.Depth(doc,Text("href"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'content' instead",len(result))
	}

	// test with tag
	s = Scrappy{}
	result = s.Depth(doc,Tag("a"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'a' instead",len(result))
	}
	s = Scrappy{}
	result = s.Depth(doc,Tag("div"))
	if len(result) != 2 {
		t.Fatal("Expected 2 nodes 'div' instead",len(result))
	}
	s = Scrappy{}
	result = s.Depth(doc,Tag("href"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'href' instead",len(result))
	}

	// test with attr
	s = Scrappy{}
	result = s.Depth(doc,Attr("href"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'href' instead",len(result))
	}
	s = Scrappy{}
	result = s.Depth(doc,Attr("onload"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'onload' instead",len(result))
	}
	s = Scrappy{}
	result = s.Depth(doc,Attr("div"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'div' instead",len(result))
	}

	// test with attr value
	s = Scrappy{}
	result = s.Depth(doc,AttrValue("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'myFunc' instead",len(result))
	}
	s = Scrappy{}
	result = s.Depth(doc,AttrValue("scrappy"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'scrappy' instead",len(result))
	}

	// mixing
	s = Scrappy{}
	result = s.Depth(doc,Attr("onload"),AttrValue("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead",len(result))
	}
	result = s.Depth(doc,Tag("a"),Attr("href"),AttrValue("scrappy"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead",len(result))
	}

	// mixing with custom filter 
	custom := func(node *html.Node) bool {
		result := s.Depth(node, Text("content"))
		return len(result) > 0
	}
	s.Depth(doc,Tag("a"),Attr("href"),AttrValue("scrappy"), custom)
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead",len(result))
	}
}

func TestScrappy_Breadth(t *testing.T) {
	s := Scrappy{}
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error",err)
	}
	// test with content
	result := s.Breadth(doc,Text("more"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'more' instead",len(result))
	}
	s = Scrappy{}
	result = s.Breadth(doc,Text("content"))
	if len(result) != 3 {
		t.Fatal("Expected 3 nodes 'content' instead",len(result))
	}
	s = Scrappy{}
	result = s.Breadth(doc,Text("href"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'content' instead",len(result))
	}

	// test with tag
	s = Scrappy{}
	result = s.Breadth(doc,Tag("a"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'a' instead",len(result))
	}
	s = Scrappy{}
	result = s.Breadth(doc,Tag("div"))
	if len(result) != 2 {
		t.Fatal("Expected 2 nodes 'div' instead",len(result))
	}
	s = Scrappy{}
	result = s.Breadth(doc,Tag("href"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'href' instead",len(result))
	}

	// test with attr
	s = Scrappy{}
	result = s.Breadth(doc,Attr("href"))
	if len(result) != 4 {
		t.Fatal("Expected 4 nodes 'href' instead",len(result))
	}
	s = Scrappy{}
	result = s.Breadth(doc,Attr("onload"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'onload' instead",len(result))
	}
	s = Scrappy{}
	result = s.Breadth(doc,Attr("div"))
	if len(result) != 0 {
		t.Fatal("Expected 0 nodes 'div' instead",len(result))
	}

	// test with attr value
	s = Scrappy{}
	result = s.Breadth(doc,AttrValue("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'myFunc' instead",len(result))
	}
	s = Scrappy{}
	result = s.Breadth(doc,AttrValue("scrappy"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes 'scrappy' instead",len(result))
	}
	// mixing
	s = Scrappy{}
	result = s.Breadth(doc,Attr("onload"),AttrValue("myFunc()"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead",len(result))
	}
	result = s.Breadth(doc,Tag("a"),Attr("href"),AttrValue("scrappy"))
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead",len(result))
	}

	// mixing with custom filter
	custom := func(node *html.Node) bool {
		result := s.Breadth(node, Text("content"))
		return len(result) > 0
	}
	s.Breadth(doc,Tag("a"),Attr("href"),AttrValue("scrappy"), custom)
	if len(result) != 1 {
		t.Fatal("Expected 1 nodes instead",len(result))
	}
}
