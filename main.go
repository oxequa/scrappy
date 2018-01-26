package scrappy

import (
	"net/http"
	"golang.org/x/net/html"
	"fmt"
	"strings"
)

type Scrappy struct {
	Nest int8
	Contains bool
	Root *html.Node
}

type FilterFunc func(node *html.Node) bool

func Tag(val string) FilterFunc{
	return func(node *html.Node) bool {return node.Type == html.ElementNode && node.Data == val }
}

func Text(val string) FilterFunc{
	return func(node *html.Node) bool {
		return node.Type == html.TextNode && strings.Contains(node.Data,val) }
}

func Attr(val string) FilterFunc{
	return func(node *html.Node) bool {
		for _, a := range node.Attr {
			if a.Key == val {
				return true
			}
		}
		return false
	}
}

func AttrValue(val string) FilterFunc{
	return func(node *html.Node) bool {
		for _, a := range node.Attr {
			if a.Val == val {
				return true
			}
		}
		return false
	}
}

// Get the content of a given url
func Get(url string) (*Scrappy, error){
	doc, err := http.Get(url)
	if err != nil{
		return nil, err
	}
	root, err := html.Parse(doc.Body)
	if err != nil{
		return nil, err
	}
	return &Scrappy{Root:root},nil
}

// Depth first search algorithm
func (s *Scrappy) Depth(node *html.Node, filters ...FilterFunc) (result []*html.Node){
	if s.Validate(node, filters...) != nil{
		result = append(result, node)
	}
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		result = append(result,s.Depth(node,filters...)...)
	}
	return
}

// Breadth first search algorithm
func (s *Scrappy) Breadth(node *html.Node, filters ...FilterFunc) (result []*html.Node){
	var breadth func(nodes []*html.Node, result []*html.Node, filters ...FilterFunc) []*html.Node
	breadth = func(nodes []*html.Node, result []*html.Node, filters ...FilterFunc) []*html.Node {
		var next []*html.Node
		for _, elm := range nodes{
			for node := elm.FirstChild; node != nil; node = node.NextSibling {
				if s.Validate(node, filters...) != nil{
					result = append(result, node)
				}
				next = append(next,node)
			}
		}
		if len(next) > 0 {
			return breadth(next, result, filters...)
		}
		return result
	}
	return breadth([]*html.Node{node},[]*html.Node{},filters...)
}

// All return all the occurrences starting from a given node
func (s *Scrappy) All(node *html.Node, filters ...FilterFunc) *[]html.Node{
	return nil
}

// All return all the occurrences starting from a given node
func (s *Scrappy) First(node *html.Node, filters ...FilterFunc) *html.Node{
	return nil
}

// All return all the occurrences starting from a given node
func (s *Scrappy) Next(node *[]html.Node, filters ...FilterFunc) *html.Node{
	return nil
}

// All return all the occurrences starting from a given node
func (s *Scrappy) Prev(node *[]html.Node, filters ...FilterFunc) *html.Node{
	return nil
}

// Validate a note by a list of filters
func (s *Scrappy) Validate(node *html.Node, filters ...FilterFunc) *html.Node{
	for _, f := range filters {
		if !f(node){
			return nil
		}
	}
	return node
}

// All return all the occurrences starting from a given node
func (s *Scrappy) Parent(node *[]html.Node, filters ...FilterFunc) *html.Node{
	return nil
}

// All return all the occurrences starting from a given node
func (s *Scrappy) Child(node *[]html.Node, filters ...FilterFunc) *[]html.Node{
	return nil
}

func (s *Scrappy) Traverse(n *html.Node, f FilterFunc) {
	for node := n.FirstChild; node != nil; node = node.NextSibling {
		fmt.Println(n.Data)
	}
}

