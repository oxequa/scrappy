package scrappy

import (
	"golang.org/x/net/html"
	"reflect"
	"strings"
	"testing"
)

const (
	weburl   = "https://github.com/oxequa"
	document = `<!DOCTYPE html>
	<html onload="myFunc()">
		<head>
			<title>title</title>
		</head>
		<body>
			body text
			<div>
				<section class="col xs" style="color:red" onclick="test()">
					<ul>
						<li>first</li>
						<li>second</li>
						<li>third</li>
						<li>forth <strong>item</strong></li>
					</ul>
					<p>a</p>
				</section>
				<section class="xs">
					<a href="">more</a>
					<a href="">content</a>
					<p>i am also a more node</p>
				</section>
				<section class="xs" style="color:red">
					<i>test</i>
				</section>
				<section>
					<i>test</i>
				</section>
			</div>
			<div>
				abc
				<strong>breadth</strong>
				<i>something</i>
				<p>more <a href="scrappy">content</a></p>
				<p><a href="">content</a> more onload</p>
			</div>
			<div>
				<img src="">
			</div>
		</body>
	</html>`
	sibling = `<ul>
		<li class="xs">1</li>
		<li class="md">2</li>
		<li class="xs">3</li>
		<li class="md">4</li>
		<li class="xs">5</li>
		<li class="md">6</li>
		<li class="xs">7</li>
		<li class="md">8</li>
	</ul>`
)

func TestNew(t *testing.T) {
	s := New()
	val := &Scrappy{}
	if reflect.TypeOf(s) != reflect.TypeOf(val) {
		t.Fatal("Unexpected type", reflect.TypeOf(s), "instead", reflect.TypeOf(val))
	}
}

func TestScrappy_Get(t *testing.T) {
	s := New()
	root, err := s.Get(weburl)
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	if root == nil {
		t.Fatal("Unexpected error root node doesn't exist")
	}
}

func TestScrappy_Nest(t *testing.T) {
	s := New()
	cp := s.Nest()
	if !cp.nested || s.nested {
		t.Fatal("Unexpected error", cp.nested, "should be true", s.nested, "should be false")
	}
}

func TestScrappy_Deep(t *testing.T) {
	s := New()
	i := 5
	cp := s.Deep(i)
	if cp.deep != i || s.deep != 0 {
		t.Fatal("Unexpected value", s.deep, "instead", i)
	}

}

func TestScrappy_Proxy(t *testing.T) {
	s := New()
	err := s.Proxy("oxequa.com:8000")
	if err != nil {
		t.Fatal(err)
	}
}

func TestScrappy_Parse(t *testing.T) {
	s := New()
	reader := strings.NewReader(document)
	_, err := s.Parse(reader)
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

}

func TestScrappy_Validate(t *testing.T) {
	s := New()
	node := &html.Node{Parent: nil, Type: html.ErrorNode}
	if s.Validate(node, Tag("div")) {
		t.Fatal("Unexpected, error nodes are invalid", node.Type)
	}
	node = &html.Node{Parent: nil, Type: html.TextNode, Data: "something"}
	if s.Validate(node) {
		t.Fatal("Unexpected, there are no filters", node.Type)
	}
	node = &html.Node{Parent: nil, Type: html.ElementNode, Data: "div"}
	if !s.Validate(node, Tag("div")) {
		t.Fatal("Unexpected, div tag is present", node.Data)
	}
	node = &html.Node{Parent: nil, Type: html.TextNode, Data: "div"}
	if s.Validate(node, Tag("div")) {
		t.Fatal("Unexpected, div tag is present but the type is wrong", node.Data, node.Type)
	}
}
