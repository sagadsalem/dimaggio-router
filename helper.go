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
