package dimaggioRouter

import (
	"net/http"
	"regexp"
)

// Handle is a function that can be registered to a route to handle HTTP
type Handle func(http.ResponseWriter, *http.Request, Params)

type route struct {
	RegexPath string
	RealPath  string
	Method    string
	Handle    Handle
	Params    []Param
}

// Router serves http
type router struct {
	routes []route
}

// NewRouter creates instance of Router
func New() *router {
	return &router{}
}

// ServeHTTP is called for every connection
func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if req.Method == route.Method {
			matched, _ := regexp.MatchString(route.RegexPath, req.URL.Path)
			if matched {
				if len(route.Params) > 0 {
					route.getParams(req.URL.Path)
				}

				route.Handle(w, req, route.Params)
				return
			}
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
	p, n := generateRegexAndParams(path)
	r.routes = append(r.routes, route{RegexPath: p, RealPath: path, Method: method, Handle: handle, Params: n})
}
