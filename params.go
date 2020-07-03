package dimaggioRouter

import "fmt"

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// Params is a Param-slice, as returned by the router.
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
