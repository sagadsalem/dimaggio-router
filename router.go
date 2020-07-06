package dimaggioRouter

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// Handle is a function that can be registered to a route to handle HTTP
type Handle func(http.ResponseWriter, *http.Request, Params)

type node struct {
	RegexPath string
	RealPath  string
	Handle    Handle
	Params    Params
}

// Router serves http
type router struct {
	routes map[string][]node

	//routes  []route
}

// NewRouter creates instance of Router
func New() *router {
	return &router{routes: make(map[string][]node)}
}

// ServeHTTP is called for every connection
func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, node := range r.routes[req.Method] {
		matched, _ := regexp.MatchString(node.RegexPath, req.URL.Path)
		if matched {
			if len(node.Params) > 0 {
				node.getParams(req.URL.Path)
			}

			node.Handle(w, req, node.Params)
			return
		}
	}
	http.NotFound(w, req)
	return
}

// GET sets get handle
func (r *router) GET(path string, handle Handle) {
	r.addRoute(http.MethodGet, path, handle)
}

// POST sets post handle
func (r *router) POST(path string, handle Handle) {
	r.addRoute(http.MethodPost, path, handle)
}

// DELETE sets delete handle
func (r *router) DELETE(path string, handle Handle) {
	r.addRoute(http.MethodDelete, path, handle)
}

// PUT sets put handle
func (r *router) PUT(path string, handle Handle) {
	r.addRoute(http.MethodPut, path, handle)
}

// add route to our routes
func (r *router) addRoute(method, path string, handle Handle) {
	r.routes[method] = append(r.routes[method], generateRoute(path, handle))
}

func generateRoute(path string, handle Handle) node {
	var s []string
	var p Params

	for index, c := range strings.Split(path, "/")[1:] {
		if strings.Contains(c, "$") {
			s = append(s, fmt.Sprint("/[a-zA-Z0-9]"))
			p = append(p, Param{
				Key:   strings.Replace(c, "$", "", -1),
				Index: index,
			})
		} else {
			s = append(s, fmt.Sprintf("/%v", c))
		}
	}
	return node{
		RegexPath: fmt.Sprintf("%v+$", strings.Join(s, "+")),
		RealPath:  path,
		Handle:    handle,
		Params:    p,
	}
}
