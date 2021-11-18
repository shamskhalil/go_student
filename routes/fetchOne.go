package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/shamskhalil/testProj/models"
)

func FetchOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	MyErr(err)
	student := models.GetOne(id)
	json.NewEncoder(w).Encode(student)
}
