package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var Address string = "localhost:3000"

func main() {
	flags := log.Ldate | log.Lshortfile
	log.SetFlags(flags)

	router := httprouter.New()

	router.GET("/students", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {})
	router.GET("/students/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {})

	err := http.ListenAndServe(Address, router)
	if err != nil {
		log.Fatal(err)
	}
}