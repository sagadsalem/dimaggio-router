package dimaggioRouter

import (
	//dimaggioRouter "github.com/sagadsalem/dimaggio-router"
	"net/http"
	"reflect"
	"testing"
)

// my testingResponseWriter that implement the ResponseWriter interface
type testingResponseWriter struct{}

// this for implementing ResponseWriter interface to use it in the Testing Router
func (trw *testingResponseWriter) Header() (h http.Header)                 { return http.Header{} }
func (trw *testingResponseWriter) Write(p []byte) (n int, err error)       { return len(p), nil }
func (trw *testingResponseWriter) WriteString(s string) (n int, err error) { return len(s), nil }
func (trw *testingResponseWriter) WriteHeader(int)                         {}

func TestNew(t *testing.T) {
	router := New()
	routed := false
	router.GET("/user/$name", func(w http.ResponseWriter, r *http.Request, dp Params) {
		routed = true
		want := Params{Param{"name", "sagad"}}
		if !reflect.DeepEqual(dp, want) {
			t.Fatalf("the values from params not matching values: want %v, got %v", want, dp)
		}
	})
	w := new(testingResponseWriter)

	req, err := http.NewRequest(http.MethodGet, "/user/sagad", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	router.ServeHTTP(w, req)

	if !routed {
		t.Fatal("routing failed in the matching phase!")
	}
}

func TestRouter_DELETE(t *testing.T) {
	router := New()
	d := false
	router.DELETE("/delete", func(w http.ResponseWriter, r *http.Request, dp Params) {
		d = true
	})

	w := new(testingResponseWriter)
	r, _ := http.NewRequest(http.MethodDelete, "/delete", nil)
	router.ServeHTTP(w, r)
	if !d {
		t.Error("delete method failed")
	}
}

func TestRouter_GET(t *testing.T) {
	router := New()
	g := false
	router.GET("/get", func(w http.ResponseWriter, r *http.Request, dp Params) {
		g = true
	})

	w := new(testingResponseWriter)
	r, _ := http.NewRequest(http.MethodGet, "/get", nil)
	router.ServeHTTP(w, r)
	if !g {
		t.Error("get method failed")
	}
}

func TestRouter_POST(t *testing.T) {
	router := New()
	p := false
	router.POST("/post", func(w http.ResponseWriter, r *http.Request, dp Params) {
		p = true
	})

	w := new(testingResponseWriter)
	r, _ := http.NewRequest(http.MethodPost, "/post", nil)
	router.ServeHTTP(w, r)
	if !p {
		t.Error("get method failed")
	}
}

func TestRouter_PUT(t *testing.T) {
	router := New()
	p := false
	router.PUT("/put", func(w http.ResponseWriter, r *http.Request, dp Params) {
		p = true
	})

	w := new(testingResponseWriter)
	r, _ := http.NewRequest(http.MethodPut, "/put", nil)
	router.ServeHTTP(w, r)
	if !p {
		t.Error("put method failed")
	}
}
