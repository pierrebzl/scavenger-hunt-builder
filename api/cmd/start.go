package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/pierrebzl/scavenger-hunt/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoomRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
