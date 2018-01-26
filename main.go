package scrappy

import (
	"golang.org/x/net/html"
	"net/http"
	"io"
)

type All struct {
	*Scrappy
}

type First struct {
	*Scrappy
}

type Scrappy struct {
	*All
	*First
	NestLevel     int
}

// FilterFunc is the general definition of a node filter
type FilterFunc func(node *html.Node) bool

func New() Scrappy{
	return Scrappy{}
}

// Validate a node by a list of filters
func Validate(node *html.Node, filters ...FilterFunc) *html.Node {
	for _, f := range filters {
		if !f(node) {
			return nil
		}
	}
	return node
}

// Get return the content of a given url
func (s *Scrappy) Get(url string) (*html.Node, error) {
	doc, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	root, err := html.Parse(doc.Body)
	if err != nil {
		return nil, err
	}
	return root, nil
}

// Parse can be used with any reader
func (s *Scrappy) Parse(reader io.ReadCloser) (*html.Node, error){
	root, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	return root, nil
}

func (a *All) Next(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.NextSibling; node != nil; node = node.NextSibling {
		if Validate(node, filters...) != nil {
			result = append(result,node)
		}
	}
	return result
}

func (a *All) Prev(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.PrevSibling; node != nil; node = node.PrevSibling {
		if Validate(node, filters...) != nil {
			result = append(result,node)
		}
	}
	return result
}

func (a *All) Child(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.FirstChild; node != nil; node = node.FirstChild {
		if Validate(node, filters...) != nil {
			result = append(result,node)
		}
	}
	return result
}

func (a *All) Depth(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	if Validate(node, filters...) != nil {
		result = append(result, node)
	}
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		result = append(result, a.Depth(node, filters...)...)
	}
	return result
}

func (a *All) Parent(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.Parent; node != nil; node = node.Parent {
		if Validate(node, filters...) != nil {
			result = append(result, node)
		}
	}
	return result
}

func (a *All) Breadth(node *html.Node, filters ...FilterFunc) []*html.Node {
	var breadth func(nodes []*html.Node, result []*html.Node, filters ...FilterFunc) []*html.Node
	breadth = func(nodes []*html.Node, result []*html.Node, filters ...FilterFunc) []*html.Node {
		var next []*html.Node
		for _, elm := range nodes {
			for node := elm.FirstChild; node != nil; node = node.NextSibling {
				if Validate(node, filters...) != nil {
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

func (f *First) Next(node *html.Node, filters ...FilterFunc) *html.Node {
	for node := node.NextSibling; node != nil; node = node.NextSibling {
		if Validate(node, filters...) != nil {
			return node
		}
	}
	return nil
}

func (f *First) Prev(node *html.Node, filters ...FilterFunc) *html.Node {
	for node := node.PrevSibling; node != nil; node = node.PrevSibling {
		if Validate(node, filters...) != nil {
			return node
		}
	}
	return nil
}

func (f *First) Child(node *html.Node, filters ...FilterFunc) *html.Node {
	for node := node.FirstChild; node != nil; node = node.FirstChild {
		if Validate(node, filters...) != nil {
			return node
		}
	}
	return nil
}

func (f *First) Depth(node *html.Node, filters ...FilterFunc) *html.Node {
	if Validate(node, filters...) != nil {
		return node
	}
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		return f.Depth(node, filters...)
	}
	return nil
}

func (f *First) Parent(node *html.Node, filters ...FilterFunc) *html.Node {
	for node := node.Parent; node != nil; node = node.Parent {
		if Validate(node, filters...) != nil {
			return node
		}
	}
	return nil
}

func (f *First) Breadth(node *html.Node, filters ...FilterFunc) *html.Node {
	var breadth func(nodes []*html.Node, filters ...FilterFunc) *html.Node
	breadth = func(nodes []*html.Node, filters ...FilterFunc) *html.Node {
		var next []*html.Node
		for _, elm := range nodes {
			for node := elm.FirstChild; node != nil; node = node.NextSibling {
				if Validate(node, filters...) != nil {
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
