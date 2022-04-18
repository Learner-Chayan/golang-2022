package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rest-go-example/pkg/controllers"
)

var LearningRoutes = func(router *mux.Router) {
	fmt.Println("")
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
}
