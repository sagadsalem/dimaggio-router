package dimaggioRouter

import (
	"fmt"
	"net/http"
	"strings"
)

// Param is a single URL parameter, consisting of a key and a value and an index
type Param struct {
	Index int
	Key   string
	Value string
}

// Params is a Param-slice
type Params []Param

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

func (ps Params) GetQuery(req *http.Request, name string) (string, error) {
	values, ok := req.URL.Query()[name]

	if !ok || len(values[0]) < 1 {
		return "", fmt.Errorf("url param '%v' is missing", name)
	}

	return values[0], nil
}

func (r *route) getParams(url string) {
	requestURL := strings.Split(url, "/")[1:]
	for i, c := range r.Params {
		r.Params[i].Value = requestURL[c.Index]
	}
}
