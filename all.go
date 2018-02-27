package scrappy

import (
	"golang.org/x/net/html"
)

// All, group of methods that return all occurrence
type A struct {
	*Scrappy
}

// Depth return nodes using first depth algorithm
func (a *A) Depth(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	validate := filters

	if a.nested && len(filters) >= 1 {
		validate = filters[:1]
	}
	if a.Validate(node, validate...) {
		if a.nested && len(filters[1:]) >= 1 {
			filters = filters[1:]
		} else {
			result = append(result, node)
		}
	}
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		result = append(result, a.Depth(node, filters...)...)
	}
	return result
}

// Breadth return nodes using first breadth algorithm
func (a *A) Breadth(node *html.Node, filters ...FilterFunc) []*html.Node {
	var breadth func(nodes []*html.Node, result []*html.Node, filters ...FilterFunc) []*html.Node
	breadth = func(nodes []*html.Node, result []*html.Node, filters ...FilterFunc) []*html.Node {
		var next []*html.Node
		validate := filters

		for _, elm := range nodes {
			for node := elm.FirstChild; node != nil; node = node.NextSibling {
				if a.nested && len(filters) >= 1 {
					validate = filters[:1]
				}
				if a.Validate(node, validate...) {
					if a.nested && len(filters[1:]) >= 1 {
						filters = filters[1:]
					} else {
						result = append(result, node)
					}
				}
				next = append(next, node)
			}
		}
		if len(next) > 0 {
			return breadth(next, result, filters...)
		}
		return result
	}
	return breadth([]*html.Node{node}, []*html.Node{}, filters...)
}

// Parent return parent nodes that matches with given filters
func (a *A) Parent(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.Parent; node != nil; node = node.Parent {
		if a.Validate(node, filters...) {
			result = append(result, node)
		}
	}
	return result
}

// Child return child nodes that matches with given filters
func (a *A) Child(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		if a.Validate(node, filters...) {
			result = append(result, node)
		}
	}
	return result
}

// Next return next sibling nodes that matches with given filters
func (a *A) NextSibling(root *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := root.NextSibling; node != nil; node = node.NextSibling {
		if node.LastChild != nil && node.PrevSibling.Data != root.Data && node.Parent != root {
			if a.Validate(node, filters...) {
				result = append(result, node)
			}
		}
	}
	return result
}

// Prev return prev sibling nodes that matches with given filters
func (a *A) PrevSibling(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.PrevSibling; node != nil; node = node.PrevSibling {
		if a.Validate(node, filters...) {
			result = append(result, node)
		}
	}
	return result
}
