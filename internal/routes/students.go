package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var students = []Student{
	{ID: "1", Name: "John Doe"},
	{ID: "2", Name: "Giga Chad"},
	{ID: "3", Name: "Chad Me Daddy"},
}

func GetStudents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	header := w.Header()

	origin := r.Header.Get("Origin")
	header.Set("Access-Control-Allow-Origin", origin)

	studentsJSON, err := json.Marshal(students)
	if err != nil {
		log.Printf("Error encoding \"students\" to json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	header.Set("Content-Type", "application/json")
	w.Write(studentsJSON)
}

func GetStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("hello from student"))
}
