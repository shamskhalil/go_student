package middlewares

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Auth(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if token := r.Header.Get("x-token"); token != "" {
			handle(w, r, p) //next()
			fmt.Println("TOKEN DETECTED ... ", token)
		} else {
			fmt.Println("TOKEN NOT FOUND ... ", token)
			http.Redirect(w, r, "/login", 302)
		}
	}
}
