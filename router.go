package dimaggioRouter

import (
	"net/http"
	"regexp"
)

// NewRouter creates instance of Router
func New() *Router {
	return &Router{
		handlers: make([]route, 0),
	}
}

// ServeHTTP is called for every connection
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.handlers {
		if r.Method == route.Method {
			matched, _ := regexp.MatchString(route.RegexPath, r.URL.Path)
			if matched {
				route.Handle(w, r, routeParams(route, r.URL.Path))
				return
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"error":"not found"}`))
	return
}

// GET sets get handle
func (router *Router) GET(path string, handle Handle) {
	router.handlers = append(router.handlers, addRoute("GET", path, handle))
}

// POST sets post handle
func (router *Router) POST(path string, handle Handle) {
	router.handlers = append(router.handlers, addRoute("POST", path, handle))
}

// DELETE sets delete handle
func (router *Router) DELETE(path string, handle Handle) {
	router.handlers = append(router.handlers, addRoute("DELETE", path, handle))
}

// PUT sets put handle
func (router *Router) PUT(path string, handle Handle) {
	router.handlers = append(router.handlers, addRoute("PUT", path, handle))
}
