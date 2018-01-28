package scrappy

import (
	"testing"
	"strings"
	"golang.org/x/net/html"
)

const (
	url      = "https://github.com/tockins"
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
				<div>
					i am also a more node
				</div>
			</p>
			<p>
				<ul>
					<li>first</li>
					<li>second</li>
					<li>third</li>
					<li>forth</li>
				</ul>
				<p>a</p>
			</p>
		</div>
		<i>something<i/>
		<p>more <a href="scrappy">content</a></p>
		<p><a href="">content</a> more onload</p>
	</body>
	</html>`
)

func TestScrappy_Get(t *testing.T) {
	s := Scrappy{}
	root, err := s.Get(url)
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	if root == nil {
		t.Fatal("Unexpected error root node doesn't exist")
	}
}

func TestScrappy_Parse(t *testing.T) {
	s := Scrappy{}
	reader := strings.NewReader(document)
	_, err := s.Parse(reader)
	if err != nil{
		t.Fatal("Unexpected error", err)
	}

}

func TestAll_Depth(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	// test with content
	result := s.All.Depth(doc, ContainText("more"))
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
	if len(result) != 2 {
		t.Fatal("Expected 2 nodes 'div' instead", len(result))
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
}

func TestAll_Breadth(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	// test with content
	result := s.All.Breadth(doc, ContainText("more"))
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
	if len(result) != 2 {
		t.Fatal("Expected 2 nodes 'div' instead", len(result))
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

func TestAll_Next(t *testing.T) {
	s := New()
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	node := s.One.Depth(doc, Tag("li"))
	result := s.All.Next(node,Tag("li"))
	if len(result) != 3 {
		t.Fatal("Expected 3 nodes 'li' instead", len(result))
	}

	node = s.One.Depth(doc, Tag("p"))
	result = s.All.Next(node,Tag("p"))
	if len(result) != 2 {
		t.Fatal("Expected 2 nodes 'p' instead", len(result))
	}
}
