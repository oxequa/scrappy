package scrappy

import (
	"golang.org/x/net/html"
	"net/http"
	"io"
)

// Scrappy is the cool struct of the lib
type Scrappy struct {
	*All
	*First
	or bool
	and bool
	deep int
}

// New return a blank scrappy instance
func New() *Scrappy{
	s := Scrappy{}
	s.All = &All{&s}
	s.First = &First{&s, 0}
	return &s
}

// Validate validate a node by a list of filters
func (s *Scrappy) Validate(node *html.Node, filters ...FilterFunc) bool {
	// check node error
	if node.Type == html.ErrorNode || len(node.Data) == 0{
		return false
	}
	// loop filters
	for _, f := range filters {
		if f(node) {
			continue
		}
		return false
	}
	return true
}

// Deep set deep option and return a new isolated scrappy
func (s *Scrappy) Deep(val int) *Scrappy{
	sc := &Scrappy{}
	sc.deep = val
	sc.All = s.All
	sc.First = s.First
	return sc
}

// Get return the content of a given url
func (s *Scrappy) Get(url string) (*html.Node, error) {
	doc, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	root, err := s.Parse(doc.Body)
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

