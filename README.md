# Dimaggio-router

<p>HTTP request router that build using go for some features not exists in net/http package mux server</p>

## Usage

```go
package main

import (
	router "github.com/sagadsalem/dimaggio-router"
	"log"
	"net/http"
)

func main() {
	r := router.New()

	r.GET("/boo", boo)
	r.GET("/get/$name", name)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func boo(w http.ResponseWriter, _ *http.Request, _ router.Params) {
	if _, err := w.Write([]byte(`{"message":"Boo Function"}`)); err != nil {
		panic(err)
	}
}

func name(w http.ResponseWriter, _ *http.Request, _ router.Params) {
	if _, err := w.Write([]byte(`{"message":"NAME Function"}`)); err != nil {
		panic(err)
	}
}

```

## Named parameters

<p>you can pass named parameters in the route by using the /$param format for example:</p>

```go
r := router.New()
r.GET("/named/parameter/$name",handler)
log.Fatal(http.ListenAndServe(":8080", r))
```

## Get parameters value

<p>so you can get parameters either by name or by index see the example below:</p>

```go
func main() {
 r := router.New()
 r.GET("/named/$param/$name",handler)
 log.Fatal(http.ListenAndServe(":8080", r))
}


func handler(w http.ResponseWriter, r *http.Request, ps router.Params) {

	param, err := ps.GetByIndex(0)
	if err != nil {
		panic(err.Error())
	}

	name, err := ps.GetByName("name")
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"message":"` + name + " - " + param + `"}`)); err != nil {
		panic(err)
	}
}
```

## Get Query String value

<p>also there are helper function to get the values of query string parameter see the example below:</p>

```go
func main() {
 r,_ := http.NewRequest("GET","/querystring?name=sagad",nil)
}


func handler(w http.ResponseWriter, r *http.Request, ps router.Params) {

	name, err := ps.GetQuery(r,"name")
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"message":"` + name + `"}`)); err != nil {
		panic(err)
	}
}
```