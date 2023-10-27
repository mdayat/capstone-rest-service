package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mdayat/capstone-rest-service/internal/routes"
)

var Address string = "localhost:3000"

func main() {
	flags := log.Ldate | log.Lshortfile
	log.SetFlags(flags)

	router := httprouter.New()

	router.GET("/students", routes.GetStudents)
	router.GET("/students/:name", routes.GetStudent)

	err := http.ListenAndServe(Address, router)
	if err != nil {
		log.Fatal(err)
	}
}
