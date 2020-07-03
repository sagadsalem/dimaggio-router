package dimaggioRouter

import (
	"net/http"
	"regexp"
)

const GET = http.MethodGet
const POST = http.MethodPost
const DELETE = http.MethodDelete
const PUT = http.MethodPut

// Handle is a function that can be registered to a route to handle HTTP
type Handle func(http.ResponseWriter, *http.Request, Params)

type route struct {
	RegexPath        string
	RealPath         string
	Method           string
	Handle           Handle
	IsNamedParameter bool
}

// Router serves http
type Router struct {
	handlers []route
	//SlashMode bool
}

// NewRouter creates instance of Router
func New() *Router {
	return &Router{
		handlers: make([]route, 0),
	}
}

// ServeHTTP is called for every connection
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.handlers {
		if req.Method == route.Method {
			matched, _ := regexp.MatchString(route.RegexPath, req.URL.Path)
			if matched {
				route.Handle(w, req, routeParams(route, req.URL.Path))
				return
			}
		}
	}
	http.NotFound(w, req)
	return
}

// GET sets get handle
func (r *Router) GET(path string, handle Handle) {
	r.addRoute(GET, path, handle)
}

// POST sets post handle
func (r *Router) POST(path string, handle Handle) {
	r.addRoute(POST, path, handle)
}

// DELETE sets delete handle
func (r *Router) DELETE(path string, handle Handle) {
	r.addRoute(DELETE, path, handle)
}

// PUT sets put handle
func (r *Router) PUT(path string, handle Handle) {
	r.addRoute(PUT, path, handle)
}

// add route to our routes
func (r *Router) addRoute(method, path string, handle Handle) {
	r.handlers = append(r.handlers, addRoute(method, path, handle))
}
