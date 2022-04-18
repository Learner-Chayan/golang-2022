package main

import (
	"github.com/gorilla/mux"
	"github.com/rest-go-example/pkg/routes"
	"log"
	"net/http"
)

func main() {
	//Create http server
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter()
	routes.LearningRoutes(router)
	log.Fatal(http.ListenAndServe(":8090", router))

}
