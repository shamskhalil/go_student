package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/shamskhalil/testProj/models"
)

func Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	MyErr(err)
	newStudent := models.Student{}
	err2 := json.NewDecoder(r.Body).Decode(&newStudent)
	MyErr(err2)
	models.UpdateOne(id, newStudent)
	json.NewEncoder(w).Encode(newStudent)
}
