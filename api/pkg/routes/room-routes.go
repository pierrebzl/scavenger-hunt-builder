package routes

import (
	"github.com/pierrebzl/scavenger-hunt/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoomRoutes = func(router *mux.Router) {
	router.HandleFunc("/room/", controllers.CreateRoom).Methods("POST")
	router.HandleFunc("/room/", controllers.GetRoom).Methods("GET")
	router.HandleFunc("/room/{id:[0-9]+}", controllers.GetRoomById).Methods("GET")
	router.HandleFunc("/room/{id:[0-9]+}", controllers.UpdateRoom).Methods("PUT")
	router.HandleFunc("/room/{id:[0-9]+}", controllers.DeleteRoom).Methods("DELETE")
}
