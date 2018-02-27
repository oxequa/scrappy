package scrappy

import (
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
)

// Scrappy is the cool struct of the lib
type Scrappy struct {
	*A
	*F
	deep   int  // deep limit
	nested bool // nested counter
}

// New return a blank scrappy instance
func New() *Scrappy {
	s := Scrappy{}
	s.A = &A{&s}
	s.F = &F{&s, 0}
	return &s
}

// Nest set nested option and return a new isolated scrappy
func (s *Scrappy) Nest() *Scrappy {
	sc := &Scrappy{}
	sc.nested = true
	sc.A = s.A
	sc.F = s.F
	return sc
}

// Deep set deep option and return a new isolated scrappy
func (s *Scrappy) Deep(val int) *Scrappy {
	sc := &Scrappy{}
	sc.deep = val
	sc.A = s.A
	sc.F = s.F
	return sc
}

// Proxy set a proxy for all requests
func (s *Scrappy) Proxy(proxy string) error {
	p, err := url.Parse(proxy)
	if err != nil {
		return err
	}
	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(p)}
	return nil
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
func (s *Scrappy) Parse(reader io.Reader) (*html.Node, error) {
	root, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	return root, nil
}

// Validate validate a node by a list of filters
func (s *Scrappy) Validate(node *html.Node, filters ...FilterFunc) bool {
	// check node error
	if node.Type == html.ErrorNode || len(node.Data) == 0 || len(filters) == 0 {
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
