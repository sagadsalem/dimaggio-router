package dimaggioRouter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParams_GetByName(t *testing.T) {
	params := Params{
		Param{Key: "id", Value: "2"},
		Param{Key: "name", Value: "sagad"},
		Param{Key: "age", Value: "23"},
		Param{Key: "type", Value: "ball"},
	}

	for i, _ := range params {
		if val, _ := params.GetByName(params[i].Key); val != params[i].Value {
			t.Errorf("Wrong value for %s: Got %s; Want %s", params[i].Key, val, params[i].Value)
		}
	}
}

func TestParams_GetByIndex(t *testing.T) {
	params := Params{
		Param{Key: "id", Value: "2"},
		Param{Key: "name", Value: "sagad"},
		Param{Key: "age", Value: "23"},
		Param{Key: "type", Value: "ball"},
	}

	for i, _ := range params {
		if val, _ := params.GetByIndex(i); val != params[i].Value {
			t.Errorf("Wrong value for %d: Got %s; Want %s", i, val, params[i].Value)
		}
	}
}

func TestParams_GetQuery(t *testing.T) {
	router := New()
	wantKey := "name"
	var value string
	router.GET("/querystring", func(w http.ResponseWriter, r *http.Request, dp Params) {
		name, err := dp.GetQuery(r, wantKey)
		if err != nil {
			t.Fatal(err.Error())
		}
		value = name
	})
	rec := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/querystring?name=ali", nil)
	router.ServeHTTP(rec, r)
	if value == "" {
		t.Fatalf("missing the query string parameter %v", wantKey)
	}
}

// benchmark testing
func BenchmarkRouter_GET(b *testing.B) {
	router := New()
	router.GET("/user/$name", benchHandler)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func benchHandler(_ http.ResponseWriter, _ *http.Request, _ Params) {}
func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	rec := httptest.NewRecorder()
	u := r.URL
	rq := u.RawQuery
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.RawQuery = rq
		router.ServeHTTP(rec, r)
	}
}
