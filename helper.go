package dimaggioRouter

import (
	"fmt"
	"strings"
)

func generateRegexAndParams(path string) (string, Params) {
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
	return fmt.Sprintf("%v+$", strings.Join(s, "+")), p
}


