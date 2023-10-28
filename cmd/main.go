package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mdayat/capstone-rest-service/internal/routes"
)

var Address string = "localhost:8080"

func main() {
	flags := log.Ldate | log.Lshortfile
	log.SetFlags(flags)

	router := httprouter.New()

	router.GET("/students", routes.GetStudents)
	router.GET("/students/:id", routes.GetStudent)

	router.POST("/auth", routes.Auth)

	err := http.ListenAndServe(Address, router)
	if err != nil {
		log.Fatal(err)
	}
}
