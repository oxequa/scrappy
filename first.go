package scrappy

import (
	"golang.org/x/net/html"
)

// First, group of methods that return only one occurrence
type First struct {
	*Scrappy
	index int
}

func (f *First) Index(index int) *First {
	f.index = index
	return f
}

// Return first child node that matches
func (f *First) Child(node *html.Node, filters ...FilterFunc) *html.Node {
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		if f.Validate(node, filters...) {
			return node
		}
	}
	return nil
}

// Return first parent node that matches
func (f *First) Parent(root *html.Node, filters ...FilterFunc) *html.Node {
	for node := root.Parent; node != nil; node = node.Parent {
		if f.Validate(node, filters...) {
			return node
		}
	}
	return nil
}

// Depth return a node using first depth algorithm, scan all nodes
func (f *First) Depth(node *html.Node, filters ...FilterFunc) *html.Node {
	if f.Validate(node, filters...) {
		return node
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
func (f *First) Breadth(node *html.Node, filters ...FilterFunc) *html.Node {
	var breadth func(nodes []*html.Node, filters ...FilterFunc) *html.Node
	breadth = func(nodes []*html.Node, filters ...FilterFunc) *html.Node {
		var next []*html.Node
		for _, elm := range nodes {
			for node := elm.FirstChild; node != nil; node = node.NextSibling {
				if f.Validate(node, filters...) {
					return node
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
func (f *First) LastChild(root *html.Node, filters ...FilterFunc) *html.Node {
	return nil
}

// LastSibling, return last sibling node that matches
func (f *First) LastSibling(root *html.Node, filters ...FilterFunc) *html.Node {
	return nil
}

// LastSibling, return initial sibling that matches
func (f *First) InitSibling(root *html.Node, filters ...FilterFunc) *html.Node {
	return nil
}

// NextSibling, return next sibling that matches
func (f *First) NextSibling(root *html.Node, filters ...FilterFunc) *html.Node {
	for node := root.NextSibling; node != nil; node = node.NextSibling {
		if len(filters) > 0 && f.Validate(node, filters...) {
			if node.LastChild != nil && node.PrevSibling.Data != root.Data && node.Parent != root {
				return node
			}
		} else {
			return node
		}
	}
	return nil
}

// PrevSibling, return prev sibling that matches
func (f *First) PrevSibling(root *html.Node, filters ...FilterFunc) *html.Node {
	for node := root.PrevSibling; node != nil; node = node.PrevSibling {
		if len(filters) > 0 && f.Validate(node, filters...) {
			if node.LastChild != nil && node.NextSibling.Data != root.Data && node.Parent != root {
				return node
			}
		} else {
			return node
		}
	}
	return nil
}

// Return i nd child node that matches
func (f *First) PositionChild(root *html.Node, i int, filters ...FilterFunc) *html.Node {
	return nil
}

// Return i nd sibling node that matches
func (f *First) PositionSibling(root *html.Node, i int, filters ...FilterFunc) *html.Node {
	return nil
}
