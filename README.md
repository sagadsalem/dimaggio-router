# Dimaggio-router

<p>HTTP request router that build using go for some features not exists in net/http package mux server</p>

## Usage

```go
package main

import (
	"github.com/sagadsalem/dimaggio-router"
	"log"
	"net/http"
)

func main() {
	r := router.New()

	r.GET("/boo", boo)
	r.GET("/get/$name", name)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func boo(w http.ResponseWriter, r *http.Request, _ router.Params) {
	if _, err := w.Write([]byte(`{"message":"Boo Function"}`)); err != nil {
		panic(err)
	}
}

func name(w http.ResponseWriter, r *http.Request, _ router.Params) {
	if _, err := w.Write([]byte(`{"message":"NAME Function"}`)); err != nil {
		panic(err)
	}
}

```