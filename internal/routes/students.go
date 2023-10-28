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
		w.Write([]byte(err.Error()))
		return
	}

	header.Set("Content-Type", "application/json")
	w.Write(studentsJSON)
}

func GetStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	header := w.Header()

	origin := r.Header.Get("Origin")
	header.Set("Access-Control-Allow-Origin", origin)

	for i, student := range students {
		if student.ID == ps.ByName("id") {
			studentJSON, err := json.Marshal(student)
			if err != nil {
				log.Printf("Error encoding \"students\" to json: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			header.Set("Content-Type", "application/json")
			w.Write(studentJSON)
			break
		} else {
			if (i + 1) == len(students) {
				errMsg := "student not found"

				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(errMsg))
			}
		}
	}
}
