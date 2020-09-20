package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/pierrebzl/scavenger-hunt/pkg/routes"
)

func NewRouter() *mux.Router {
    r := mux.NewRouter()
    routes.RoomRoutes(r)
    routes.SpotRoutes(r)
    return r
}

func main() {
	r := NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
