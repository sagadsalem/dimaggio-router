# Dimaggio-router [![Build Status](https://api.travis-ci.org/sagadsalem/dimaggio-router.svg?branch=master&status=passed)](https://travis-ci.org/github/sagadsalem/dimaggio-router) 


> HTTP request router that builds using go for some features does not exist in net/http package mux server

## Installation

```shell script
go get -u github.com/sagadsalem/dimaggio-router
```

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
	log.Fatal(http.ListenAndServe(":8080", r))
}

func boo(w http.ResponseWriter, _ *http.Request, _ router.Params) {
	if _, err := w.Write([]byte(`{"message":"Boo Function"}`)); err != nil {
		panic(err)
	}
}

```


## Named parameter

you can pass named parameters in the route by using the [$param] format and you can get parameters either by name or by index see the example below:

```go
package main

import (
	router "github.com/sagadsalem/dimaggio-router"
	"log"
	"net/http"
)

func main() {
    r := router.New()
    r.GET("/get/user/$id",handler)
    log.Fatal(http.ListenAndServe(":8080", r))
}

func handler(w http.ResponseWriter, r *http.Request, dp router.Params) {
	param, err := dp.GetByIndex(0)
	if err != nil {
		panic(err.Error())
	}

	id, err := dp.GetByName("id")
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"message":"` + id + " - " + param + `"}`)); err != nil {
		panic(err)
	}
}
```

## Query String

also, there is a helper function to get the values of the query string parameter to see the example below:

```go
package main

import (
    "net/http"
	router "github.com/sagadsalem/dimaggio-router"
)

func main() {
    r,_ := http.NewRequest("GET","/querystring?name=sagad",nil)
}

func handler(w http.ResponseWriter, r *http.Request, dp router.Params) {
	name, err := dp.GetQuery(r,"name")
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"message":"` + name + `"}`)); err != nil {
		panic(err)
	}
}
```

## TODO

- [x] Named parameter
- [x] Querystring parameter
- [ ] Middleware
- [ ] Route Group
- [ ] Enable CORS