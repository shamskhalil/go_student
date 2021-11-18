package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shamskhalil/testProj/models"
)

func FetchAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	allStudents := models.GetAll()
	json.NewEncoder(w).Encode(allStudents)
}
