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

func (ps Params) GetByName(name string) (string, error) {
	for _, p := range ps {
		if p.Key == name {
			return p.Value, nil
		}
	}
	return "", fmt.Errorf("the parameter %v was not found in the request", name)
}

func (ps Params) GetByIndex(index int) (string, error) {
	for i, p := range ps {
		if i == index {
			return p.Value, nil
		}
	}
	return "", fmt.Errorf("the index %d was not found in the request", index)
}

/*
* Helpers functions
* addRoute function return a route with all the required fields
* routePath function return the path in regex format for faster matching in the routing
 */
func addRoute(method string, path string, handle Handle) route {
	p, n := routePath(path)
	return route{
		RegexPath:        p,
		RealPath:         path,
		Method:           method,
		Handle:           handle,
		IsNamedParameter: n,
	}
}

func routePath(path string) (string, bool) {
	p := ""
	var isNamedParameter = false
	var s []string

	components := strings.Split(path, "/")[1:]
	for _, c := range components {
		if strings.Contains(c, "$") {
			s = append(s, fmt.Sprint("/[a-zA-Z0-9]"))
			isNamedParameter = true
		} else {
			s = append(s, fmt.Sprintf("/%v", c))
		}
	}
	p = strings.Join(s, "+") + "+$"
	return p, isNamedParameter
}

func routeParams(route route, url string) Params {
	if route.IsNamedParameter == true {
		var params Params
		realComponents := strings.Split(route.RealPath, "/")[1:]
		urlComponents := strings.Split(url, "/")[1:]
		for index, c := range realComponents {
			if strings.Contains(c, "$") {
				params = append(params, Param{
					Key:   strings.Replace(c, "$", "", -1), // without the $ sign
					Value: urlComponents[index],
				})
			}
		}
		return params
	}
	return Params{}
}
