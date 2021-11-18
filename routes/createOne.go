package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shamskhalil/testProj/models"
)

func CreateOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	newStudent := models.Student{}
	err2 := json.NewDecoder(r.Body).Decode(&newStudent)
	MyErr(err2)
	models.AddOne(newStudent)
	json.NewEncoder(w).Encode(newStudent)
}
