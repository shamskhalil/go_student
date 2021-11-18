package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/shamskhalil/testProj/models"
)

func Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	MyErr(err)
	deletedStudent := models.GetOne(id)
	models.DeleteOne(id)
	json.NewEncoder(w).Encode(deletedStudent)
}
