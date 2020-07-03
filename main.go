package main

import (
	"github.com/sagadsalem/dimaggio-router"
	"log"
	"net/http"
)

func main() {

	r := dimaggioRouter.New()
	r.GET("/boo", boo)

	r.GET("/foo/bar", fooBar)
	r.GET("/get/$name", name)

	r.GET("/get/$name/after", getNameAfter)
	r.POST("/get/$name/after", postNameAfter)

	r.GET("/get/$name/$param", paramAfter)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func boo(w http.ResponseWriter, r *http.Request, _ dimaggioRouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"message":"Boo"}`)); err != nil {
		panic(err)
	}
}
func fooBar(w http.ResponseWriter, r *http.Request, _ dimaggioRouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"message":"fooBar"}`)); err != nil {
		panic(err)
	}
}
func name(w http.ResponseWriter, r *http.Request, ps dimaggioRouter.Params) {
	name, err := ps.GetByName("name")
	if err != nil {
		panic(err.Error())
	}

	//index, err := ps.GetByIndex(0)
	//if err != nil {
	//	panic(err.Error())
	//}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"message":"` + name + `"}`)); err != nil {
		panic(err)
	}
}

func getNameAfter(w http.ResponseWriter, r *http.Request, _ dimaggioRouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"message":"GETNAMEAFTER"}`)); err != nil {
		panic(err)
	}
}
func postNameAfter(w http.ResponseWriter, r *http.Request, _ dimaggioRouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"message":"POSTNAMEAFTER"}`)); err != nil {
		panic(err)
	}
}

func paramAfter(w http.ResponseWriter, r *http.Request, ps dimaggioRouter.Params) {

	param, err := ps.GetByName("param")
	if err != nil {
		panic(err.Error())
	}

	name, err := ps.GetByName("name")
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"message":"` + name + " " + param + `"}`)); err != nil {
		panic(err)
	}
}
