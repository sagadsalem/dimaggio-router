package dimaggioRouter_test

import (
	dimaggioRouter "github.com/sagadsalem/dimaggio-router"
	"net/http"
	"testing"
)

func TestParams_GetByName(t *testing.T) {
	params := dimaggioRouter.Params{
		dimaggioRouter.Param{Key: "id", Value: "2"},
		dimaggioRouter.Param{Key: "name", Value: "sagad"},
		dimaggioRouter.Param{Key: "age", Value: "23"},
		dimaggioRouter.Param{Key: "type", Value: "ball"},
	}

	for i, _ := range params {
		if val, _ := params.GetByName(params[i].Key); val != params[i].Value {
			t.Errorf("Wrong value for %s: Got %s; Want %s", params[i].Key, val, params[i].Value)
		}
	}
}

func TestParams_GetByIndex(t *testing.T) {
	params := dimaggioRouter.Params{
		dimaggioRouter.Param{Key: "id", Value: "2"},
		dimaggioRouter.Param{Key: "name", Value: "sagad"},
		dimaggioRouter.Param{Key: "age", Value: "23"},
		dimaggioRouter.Param{Key: "type", Value: "ball"},
	}

	for i, _ := range params {
		if val, _ := params.GetByIndex(i); val != params[i].Value {
			t.Errorf("Wrong value for %d: Got %s; Want %s", i, val, params[i].Value)
		}
	}
}

func TestParams_GetQuery(t *testing.T) {
	router := dimaggioRouter.New()
	wantKey := "name"
	var value string
	router.GET("/querystring", func(w http.ResponseWriter, r *http.Request, dp dimaggioRouter.Params) {
		name, err := dp.GetQuery(r, wantKey)
		if err != nil {
			t.Fatal(err.Error())
		}
		value = name
	})
	w := new(testingResponseWriter)
	r, _ := http.NewRequest(http.MethodGet, "/querystring?name=ali", nil)
	router.ServeHTTP(w, r)
	if value == "" {
		t.Fatalf("missing the query string parameter %v", wantKey)
	}
}
