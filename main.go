package scrappy

import (
	"golang.org/x/net/html"
	"net/http"
	"io"
	"github.com/go-siris/siris/core/router/node"
)

// All, group of methods that return all occurrence
type All struct {
	*Scrappy
}

// One, group of methods that return only one occurrence
type One struct {
	*Scrappy
}

// Scrappy is the cool struct of the lib
type Scrappy struct {
	*All
	*One
	deep int
	index int
}

// New return a blank scrappy instance
func New() Scrappy{
	s := Scrappy{}
	s.All = &All{&s}
	s.One = &One{&s}
	return s
}

// Next return next sibling nodes that matches with given filters
func (a *All) Next(root *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := root.NextSibling; node != nil; node = node.NextSibling {
		if node.LastChild != nil && node.PrevSibling.Data != root.Data && node.Parent != root {
			if a.Validate(node, filters...) != nil {
				result = append(result, node)
			}
		}
	}
	return result
}

// Prev return prev sibling nodes that matches with given filters
func (a *All) Prev(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.PrevSibling; node != nil; node = node.PrevSibling {
		if a.Validate(node, filters...) != nil {
			result = append(result,node)
		}
	}
	return result
}

// Child return child nodes that matches with given filters
func (a *All) Child(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.FirstChild; node != nil; node = node.FirstChild {
		if a.Validate(node, filters...) != nil {
			result = append(result,node)
		}
	}
	return result
}

// Depth return nodes using first depth algorithm
func (a *All) Depth(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	if a.Validate(node, filters...) != nil {
		result = append(result, node)
	}
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		result = append(result, a.Depth(node, filters...)...)
	}
	return result
}

// Parent return parent nodes that matches with given filters
func (a *All) Parent(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.Parent; node != nil; node = node.Parent {
		if a.Validate(node, filters...) != nil {
			result = append(result, node)
		}
	}
	return result
}

// Breadth return nodes using first breadth algorithm
func (a *All) Breadth(node *html.Node, filters ...FilterFunc) []*html.Node {
	var breadth func(nodes []*html.Node, result []*html.Node, filters ...FilterFunc) []*html.Node
	breadth = func(nodes []*html.Node, result []*html.Node, filters ...FilterFunc) []*html.Node {
		var next []*html.Node
		for _, elm := range nodes {
			for node := elm.FirstChild; node != nil; node = node.NextSibling {
				if a.Validate(node, filters...) != nil {
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

// Next return a next sibling node that matches with given filters
func (o *One) Next(root *html.Node, filters ...FilterFunc) *html.Node {
	for node := root.NextSibling; node != nil; node = node.NextSibling {
		if node.LastChild != nil && node.PrevSibling.Data != root.Data && node.Parent != root {
			if o.Validate(node, filters...) != nil {
				return node
			}
		}
	}
	return nil
}

// Prev return a prev sibling node that matches with given filters
func (o *One) Prev(node *html.Node, filters ...FilterFunc) *html.Node {
	for node := node.PrevSibling; node != nil; node = node.PrevSibling {
		if o.Validate(node, filters...) != nil {
			return node
		}
	}
	return nil
}

// Child return a child node that matches with given filters
func (o *One) Child(node *html.Node, filters ...FilterFunc) *html.Node {
	for node := node.FirstChild; node != nil; node = node.FirstChild {
		if o.Validate(node, filters...) != nil {
			return node
		}
	}
	return nil
}

// Depth return a node using first depth algorithm
func (o *One) Depth(node *html.Node, filters ...FilterFunc) *html.Node {
	if o.Validate(node, filters...) != nil {
		return node
	}
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		result := o.Depth(node, filters...)
		if result != nil{
			return result
		}
	}
	return nil
}

// Parent return a parent node that matches with given filters
func (o *One) Parent(node *html.Node, filters ...FilterFunc) *html.Node {
	for node := node.Parent; node != nil; node = node.Parent {
		if o.Validate(node, filters...) != nil {
			return node
		}
	}
	return nil
}

// Breadth return a node using first breadth algorithm
func (o *One) Breadth(node *html.Node, filters ...FilterFunc) *html.Node {
	var breadth func(nodes []*html.Node, filters ...FilterFunc) *html.Node
	breadth = func(nodes []*html.Node, filters ...FilterFunc) *html.Node {
		var next []*html.Node
		for _, elm := range nodes {
			for node := elm.FirstChild; node != nil; node = node.NextSibling {
				if o.Validate(node, filters...) != nil {
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

// Deep set deep option and return a new isolated scrappy
func (s *Scrappy) Deep(val int) *Scrappy{
	return &Scrappy{s.All,s.One,val,0}
}

// Deep set index option and return a new isolated scrappy
func (s *Scrappy) Index(val int) *Scrappy{
	return &Scrappy{s.All,s.One,0,val}
}

// Mix set all the options and return a new isolated scrappy
func (s *Scrappy) Mix(index int, deep int) *Scrappy{
	return &Scrappy{s.All,s.One,deep,index}
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
func (s *Scrappy) Parse(reader io.Reader) (*html.Node, error){
	root, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	return root, nil
}

// Validate validate a node by a list of filters
func (s *Scrappy) Validate(node *html.Node, filters ...FilterFunc) *html.Node {
	// check node error
	if node.Type == html.ErrorNode{
		return nil
	}
	// loop filters
	for _, f := range filters {
		if f(node) {
			continue
		}
		return nil
	}
	// check for an index
	if s.index != 0{
		s.index--
		return nil
	}
	return node
}
