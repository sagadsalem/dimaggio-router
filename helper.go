package dimaggioRouter

import (
	"fmt"
	"strings"
)

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
	var isNamedParameter = false
	var s []string

	for _, c := range strings.Split(path, "/")[1:] {
		if strings.Contains(c, "$") {
			s = append(s, fmt.Sprint("/[a-zA-Z0-9]"))
			isNamedParameter = true
		} else {
			s = append(s, fmt.Sprintf("/%v", c))
		}
	}
	return strings.Join(s, "+") + "+$", isNamedParameter
}

func routeParams(route route, url string) Params {
	if route.IsNamedParameter == true {
		var params Params
		urlComponents := strings.Split(url, "/")[1:]
		for index, c := range strings.Split(route.RealPath, "/")[1:] {
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
