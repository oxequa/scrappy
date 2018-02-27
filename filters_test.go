package scrappy

import (
	"golang.org/x/net/html"
	"testing"
)

func TestTag(t *testing.T) {
	node := &html.Node{
		Type: html.ElementNode,
		Data: "test",
		Attr: []html.Attribute{
			{
				Key: "TestKey",
				Val: "TestVal",
			},
		},
	}
	f := Tag("test")
	if !f(node) {
		t.Fatal("Unexpected error, given an element node")
	}
	node.Type = html.TextNode
	if f(node) {
		t.Fatal("Unexpected error, given a text node")
	}
	f = Tag("testing")
	if f(node) {
		t.Fatal("Unexpected error")
	}
}

func TestText(t *testing.T) {
	node := &html.Node{
		Type: html.TextNode,
		Data: "test",
		Attr: []html.Attribute{
			{
				Key: "TestKey",
				Val: "TestVal",
			},
		},
	}
	f := Text("test")
	if !f(node) {
		t.Fatal("Unexpected error, given an element node")
	}
	node.Type = html.DocumentNode
	if f(node) {
		t.Fatal("Unexpected error, given a documnet node")
	}
	f = Text("testing")
	if f(node) {
		t.Fatal("Unexpected error")
	}
}

func TestAttr(t *testing.T) {
	node := &html.Node{
		Type: html.ElementNode,
		Data: "test",
		Attr: []html.Attribute{
			{
				Key: "TestKey",
				Val: "TestVal",
			},
			{
				Key: "TestKeyB",
				Val: "TestValB",
			},
		},
	}
	f := Attr("TestKey")
	if !f(node) {
		t.Fatal("Unexpected error")
	}
	f = Attr("TestKey1")
	if f(node) {
		t.Fatal("Unexpected error")
	}
	f = Attr("TestKeyB")
	if !f(node) {
		t.Fatal("Unexpected error")
	}
}

func TestValue(t *testing.T) {
	node := &html.Node{
		Type: html.ElementNode,
		Data: "test",
		Attr: []html.Attribute{
			{
				Key: "TestKey",
				Val: "TestVal",
			},
			{
				Key: "TestKeyB",
				Val: "TestValB",
			},
		},
	}
	f := Value("TestVal")
	if !f(node) {
		t.Fatal("Unexpected error")
	}
	f = Value("TestVal1")
	if f(node) {
		t.Fatal("Unexpected error")
	}
	f = Value("TestValB")
	if !f(node) {
		t.Fatal("Unexpected error")
	}
}

func TestAttrVal(t *testing.T) {
	node := &html.Node{
		Type: html.ElementNode,
		Data: "test",
		Attr: []html.Attribute{
			{
				Key: "TestKey",
				Val: "TestVal",
			},
			{
				Key: "TestKeyB",
				Val: "TestValB",
			},
		},
	}
	f := AttrVal("TestKey", "TestVal")
	if !f(node) {
		t.Fatal("Unexpected error")
	}
	f = AttrVal("TestKey1", "TestValB")
	if f(node) {
		t.Fatal("Unexpected error")
	}
	f = AttrVal("TestKeyB", "TestValB")
	if !f(node) {
		t.Fatal("Unexpected error")
	}
}

func TestContainsTag(t *testing.T) {
	node := &html.Node{
		Type: html.ElementNode,
		Data: "test",
		Attr: []html.Attribute{
			{
				Key: "TestKey",
				Val: "TestVal",
			},
		},
	}
	f := ContainsTag("tes")
	if !f(node) {
		t.Fatal("Unexpected error")
	}
	f = ContainsTag("test1")
	if f(node) {
		t.Fatal("Unexpected error")
	}
}

func TestContainsText(t *testing.T) {
	node := &html.Node{
		Type: html.TextNode,
		Data: "test",
		Attr: []html.Attribute{
			{
				Key: "TestKey",
				Val: "TestVal",
			},
		},
	}
	f := ContainsText("tes")
	if !f(node) {
		t.Fatal("Unexpected error")
	}
	f = ContainsText("test1")
	if f(node) {
		t.Fatal("Unexpected error")
	}
}

func TestContainsAttr(t *testing.T) {
	node := &html.Node{
		Type: html.TextNode,
		Data: "test",
		Attr: []html.Attribute{
			{
				Key: "TestKey",
				Val: "TestVal",
			},
		},
	}
	f := ContainsAttr("est")
	if !f(node) {
		t.Fatal("Unexpected error")
	}
	f = ContainsAttr("key1")
	if f(node) {
		t.Fatal("Unexpected error")
	}
}

func TestContainsValue(t *testing.T) {
	node := &html.Node{
		Type: html.TextNode,
		Data: "test",
		Attr: []html.Attribute{
			{
				Key: "TestKey",
				Val: "TestVal",
			},
		},
	}
	f := ContainsValue("est")
	if !f(node) {
		t.Fatal("Unexpected error")
	}
	f = ContainsValue("val")
	if f(node) {
		t.Fatal("Unexpected error")
	}
}
