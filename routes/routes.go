package routes

import (
	"api-rest-go/controllers"
	"log"
	"net/http"
)

func HandlerRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
