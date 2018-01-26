package scrappy

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

type Scrappy struct {
	first         bool
	Contains      bool
	DepthSearch   bool
	BreadthSearch bool
	NestLevel     int
	Root          *html.Node
}

// FilterFunc is the general definition of a node filter
type FilterFunc func(node *html.Node) bool

// Tag  is a filter func that return a node that matches with a given tag
func Tag(val string) FilterFunc {
	return func(node *html.Node) bool { return node.Type == html.ElementNode && node.Data == val }
}

// Text  is a filter func that return a node that matches with a given text content
func Text(val string) FilterFunc {
	return func(node *html.Node) bool {
		return node.Type == html.TextNode && strings.Contains(node.Data, val)
	}
}

// Attr is a filter func that return a node that matches with a given attribute
func Attr(val string) FilterFunc {
	return func(node *html.Node) bool {
		for _, a := range node.Attr {
			if a.Key == val {
				return true
			}
		}
		return false
	}
}

// AttrValues is a filter func that return a node that matches with a given attribute value
func AttrValue(val string) FilterFunc {
	return func(node *html.Node) bool {
		for _, a := range node.Attr {
			if a.Val == val {
				return true
			}
		}
		return false
	}
}

// Get return the content of a given url
func Get(url string) (*Scrappy, error) {
	doc, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	root, err := html.Parse(doc.Body)
	if err != nil {
		return nil, err
	}
	return &Scrappy{Root: root, DepthSearch: true}, nil
}

// Depth first search algorithm
func (s *Scrappy) Depth(node *html.Node, filters ...FilterFunc) (result []*html.Node) {
	if s.first && len(result) == 1 {
		return
	}
	if s.Validate(node, filters...) != nil {
		result = append(result, node)
	}
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		result = append(result, s.Depth(node, filters...)...)
	}
	return
}

// Breadth first search algorithm
func (s *Scrappy) Breadth(node *html.Node, filters ...FilterFunc) (result []*html.Node) {
	var breadth func(nodes []*html.Node, result []*html.Node, filters ...FilterFunc) []*html.Node
	breadth = func(nodes []*html.Node, result []*html.Node, filters ...FilterFunc) []*html.Node {
		var next []*html.Node
		for _, elm := range nodes {
			for node := elm.FirstChild; node != nil; node = node.NextSibling {
				if s.first && len(result) == 1 {
					return result
				}
				if s.Validate(node, filters...) != nil {
					result = append(result, node)
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

// All return all the occurrences starting from a given node
func (s *Scrappy) All(node *html.Node, filters ...FilterFunc) []*html.Node {
	if s.BreadthSearch {
		return s.Breadth(node, filters...)
	}
	return s.Depth(node, filters...)
}

// First return first occurrence starting from a given node
func (s *Scrappy) First(node *html.Node, filters ...FilterFunc) *html.Node {
	var result []*html.Node
	s.first = true
	if s.BreadthSearch {
		result = s.Breadth(node, filters...)
	}
	result = s.Depth(node, filters...)
	s.first = false
	if len(result) > 0 {
		return result[0]
	}
	return nil
}

// Next return
func (s *Scrappy) Next(node *[]html.Node, filters ...FilterFunc) *html.Node {
	return nil
}

// All return all the occurrences starting from a given node
func (s *Scrappy) Prev(node *[]html.Node, filters ...FilterFunc) *html.Node {
	return nil
}

// Validate a note by a list of filters
func (s *Scrappy) Validate(node *html.Node, filters ...FilterFunc) *html.Node {
	for _, f := range filters {
		if !f(node) {
			return nil
		}
	}
	return node
}

// All return all the occurrences starting from a given node
func (s *Scrappy) Parent(node *[]html.Node, filters ...FilterFunc) *html.Node {
	return nil
}

// All return all the occurrences starting from a given node
func (s *Scrappy) Child(node *[]html.Node, filters ...FilterFunc) *[]html.Node {
	return nil
}

func (s *Scrappy) Traverse(n *html.Node, f FilterFunc) {
	for node := n.FirstChild; node != nil; node = node.NextSibling {
		fmt.Println(n.Data)
	}
}
