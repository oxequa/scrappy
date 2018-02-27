package scrappy

import (
	"golang.org/x/net/html"
)

// First, group of methods that return only one occurrence
type F struct {
	*Scrappy
	index int
}

func decrease(i int) int{
	if i > 0 {
		i--
	}
	return i
}

func (f *F) Index(index int) *F {
	f.index = index
	return f
}

// Return first parent node that matches
func (f *F) Parent(root *html.Node, filters ...FilterFunc) *html.Node {
	for node := root.Parent; node != nil; node = node.Parent {
		if f.Validate(node, filters...) {
			if f.index = decrease(f.index); f.index == 0 {
				return node
			}
		}
	}
	return nil
}

// Depth return a node using first depth algorithm, scan all nodes
func (f *F) Depth(node *html.Node, filters ...FilterFunc) *html.Node {
	if f.Validate(node, filters...) {
		if f.index = decrease(f.index); f.index == 0 {
			return node
		}
	}
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		result := f.Depth(node, filters...)
		if result != nil {
			return result
		}
	}
	return nil
}

// Breadth return a node using first breadth algorithm, scan all nodes
func (f *F) Breadth(node *html.Node, filters ...FilterFunc) *html.Node {
	var breadth func(nodes []*html.Node, filters ...FilterFunc) *html.Node
	breadth = func(nodes []*html.Node, filters ...FilterFunc) *html.Node {
		var next []*html.Node
		for _, elm := range nodes {
			for node := elm.FirstChild; node != nil; node = node.NextSibling {
				if f.Validate(node, filters...) {
					if f.index = decrease(f.index); f.index == 0 {
						return node
					}
				}
				next = append(next, node)
			}
		}
		if len(next) > 0 {
			return breadth(next, filters...)
		}
		return nil
	}
	return breadth([]*html.Node{node}, filters...)
}

// Return last child, it's like a last sibling with the first child
func (f *F) LastChild(node *html.Node, filters ...FilterFunc) *html.Node {
	for node := node.LastChild; node != nil; node = node.PrevSibling {
		if f.Validate(node, filters...) {
			if f.index = decrease(f.index); f.index == 0 {
				return node
			}
		}
	}
	return nil
}

// Return first child node that matches
func (f *F) FirstChild(node *html.Node, filters ...FilterFunc) *html.Node {
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		if f.Validate(node, filters...) {
			if f.index = decrease(f.index); f.index == 0 {
				return node
			}
		}
	}
	return nil
}

// LastSibling, return last sibling node that matches
func (f *F) LastSibling(node *html.Node, filters ...FilterFunc) (result *html.Node) {
	for node := node.NextSibling; node != nil; node = node.NextSibling {
		if f.Validate(node, filters...) {
			if f.index = decrease(f.index); f.index == 0 {
				result = node
			}
		}
	}
	return result
}

// FirstSibling, return first sibling node that matches
func (f *F) FirstSibling(node *html.Node, filters ...FilterFunc) (result *html.Node) {
	for node := node.PrevSibling; node != nil; node = node.PrevSibling {
		if f.Validate(node, filters...) {
			if f.index = decrease(f.index); f.index == 0 {
				result = node
			}
		}
	}
	return result
}

// NextSibling, return next sibling that matches
func (f *F) NextSibling(node *html.Node, filters ...FilterFunc) *html.Node {
	for node := node.NextSibling; node != nil; node = node.NextSibling {
		if f.Validate(node, filters...) {
			if f.index = decrease(f.index); f.index == 0 {
				return node
			}
		}
	}
	return nil
}

// PrevSibling, return prev sibling that matches
func (f *F) PrevSibling(node *html.Node, filters ...FilterFunc) *html.Node {
	for node := node.PrevSibling; node != nil; node = node.PrevSibling {
		if f.Validate(node, filters...) {
			if f.index = decrease(f.index); f.index == 0 {
				return node
			}
		}
	}
	return nil
}