package routes

import (
	"api-rest-go/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.GetPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
