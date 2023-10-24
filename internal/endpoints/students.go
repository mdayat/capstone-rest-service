package endpoints

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Student struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

func GetStudents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("hello from students"))
}

func GetStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("hello from student"))
}