package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shamskhalil/testProj/middlewares"
	"github.com/shamskhalil/testProj/models"
	"github.com/shamskhalil/testProj/routes"
)

func main() {
	models.AddOne(models.Student{FirstName: "Khalil", LastName: "M. Shams", Age: 22})
	models.AddOne(models.Student{FirstName: "Jamilu", LastName: "Sani", Age: 43})
	models.AddOne(models.Student{FirstName: "Amina", LastName: "Salis", Age: 19})

	router := httprouter.New()

	router.GET("/student", middlewares.JHeader(routes.FetchAll))
	router.POST("/student", middlewares.JHeader(middlewares.Auth(routes.CreateOne)))
	router.GET("/student/:id", middlewares.JHeader(routes.FetchOne))
	router.PUT("/student/:id", middlewares.JHeader(routes.Update))
	router.DELETE("/student/:id", middlewares.JHeader(middlewares.Auth(routes.Delete)))
	router.GET("/login", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Login Page")
	})
	router.NotFound = http.FileServer(http.Dir("./www"))

	fmt.Println("Server listening on port 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
