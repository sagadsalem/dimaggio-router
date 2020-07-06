package dimaggioRouter

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	router := New()
	router.GET("/user/$name", func(w http.ResponseWriter, r *http.Request, dp Params) {
		want := Params{Param{1, "name", "sagad"}}
		if !reflect.DeepEqual(dp, want) {
			t.Fatalf("the values from params not matching values: want %v, got %v", want, dp)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("checked")); err != nil {
			t.Fatalf(err.Error())
		}
	})

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/user/sagad", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	//routing...
	router.ServeHTTP(rec, req)

	// get result
	resp := rec.Result()
	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("excpected status OK; got %v", resp.StatusCode)
	}

	// read the result body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert to string remove all white spaces
	bodyString := strings.TrimSpace(string(bodyBytes))

	// check if body equals to checked
	if bodyString != "checked" {
		t.Fatalf("excpected response of checked; got %v", bodyString)
	}
}

func TestRouter_DELETE(t *testing.T) {
	// init the router
	router := New()
	router.DELETE("/delete", func(w http.ResponseWriter, r *http.Request, dp Params) {
		if _, err := w.Write([]byte("delete")); err != nil {
			t.Fatal(err.Error())
		}
	})

	// create recorder and request
	rec := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodDelete, "/delete", nil)

	// routing...
	router.ServeHTTP(rec, r)

	// get result
	resp := rec.Result()
	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("excpected status OK; got %v", resp.StatusCode)
	}
	// read the result body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert to string remove all white spaces
	bodyString := strings.TrimSpace(string(bodyBytes))

	// check if body equals to delete
	if bodyString != "delete" {
		t.Fatalf("excpected response of delete; got %v", bodyString)
	}
}

func TestRouter_GET(t *testing.T) {
	// init the router
	router := New()
	router.GET("/get", func(w http.ResponseWriter, r *http.Request, dp Params) {
		if _, err := w.Write([]byte("get")); err != nil {
			t.Fatal(err.Error())
		}
	})

	// create recorder and request
	rec := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/get", nil)

	// routing...
	router.ServeHTTP(rec, r)

	// get result
	resp := rec.Result()
	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("excpected status OK; got %v", resp.StatusCode)
	}
	// read the result body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert to string remove all white spaces
	bodyString := strings.TrimSpace(string(bodyBytes))

	// check if body equals to get
	if bodyString != "get" {
		t.Fatalf("excpected response of get; got %v", bodyString)
	}
}

func TestRouter_POST(t *testing.T) {
	// init the router
	router := New()
	router.POST("/post", func(w http.ResponseWriter, r *http.Request, dp Params) {
		if _, err := w.Write([]byte("post")); err != nil {
			t.Fatal(err.Error())
		}
	})

	// create recorder and request
	rec := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPost, "/post", nil)

	// routing...
	router.ServeHTTP(rec, r)

	// get result
	resp := rec.Result()
	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("excpected status OK; got %v", resp.StatusCode)
	}

	// read the result body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert to string remove all white spaces
	bodyString := strings.TrimSpace(string(bodyBytes))

	// check if body equals to post
	if bodyString != "post" {
		t.Fatalf("excpected response of post; got %v", bodyString)
	}
}

func TestRouter_PUT(t *testing.T) {
	// init the router
	router := New()
	router.PUT("/put", func(w http.ResponseWriter, r *http.Request, dp Params) {
		if _, err := w.Write([]byte("put")); err != nil {
			t.Fatal(err.Error())
		}
	})

	// create recorder and request
	rec := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPut, "/put", nil)

	// routing...
	router.ServeHTTP(rec, r)

	// get result
	resp := rec.Result()
	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("excpected status OK; got %v", resp.StatusCode)
	}

	// read the result body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert to string remove all white spaces
	bodyString := strings.TrimSpace(string(bodyBytes))

	// check if body equals to put
	if bodyString != "put" {
		t.Fatalf("excpected response of put; got %v", bodyString)
	}
}
