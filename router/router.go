package router

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// Handle is a function that can be registered to a route to handle HTTP
type Handle func(http.ResponseWriter, *http.Request, Params)

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// Params is a Param-slice, as returned by the router.
type Params []Param

type route struct {
	Path   string
	Method string
	Handle Handle
}

// Router serves http
type Router struct {
	handlers []route
	//SlashMode bool
}

// NewRouter creates instance of Router
func New() *Router {
	//s := false
	//if len(slashMode) > 0 && len(slashMode) < 2 {
	//	s = slashMode[0]
	//}
	return &Router{
		handlers: make([]route, 0),
		//SlashMode: s,
	}
}

// ServeHTTP is called for every connection
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.handlers {
		if r.Method == route.Method {
			matched, _ := regexp.MatchString(route.Path, r.URL.Path)
			if matched {
				route.Handle(w, r, Params{})
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

/*
* Helpers functions
* addRoute function return a route with all the required fields
* routePath function return the path in regex format for faster matching in the routing
 */
func addRoute(method string, path string, handle Handle) route {
	return route{
		Path:   routePath(path),
		Method: method,
		Handle: handle,
	}
}

func routePath(path string) string {
	components := strings.Split(path, "/")[1:]
	p := ""
	var s []string
	for _, c := range components {
		if strings.Contains(c, "$") {
			s = append(s, fmt.Sprint("/[a-zA-Z0-9]"))
		} else {
			s = append(s, fmt.Sprintf("/%v", c))
		}
	}
	p = strings.Join(s, "+") + "+$"
	return p
}
