package routes

import (
	"github.com/pierrebzl/scavenger-hunt/pkg/controllers"
	"github.com/gorilla/mux"
)

var SpotRoutes = func(router *mux.Router) {
	router.HandleFunc("/room/{roomId:[0-9]+}/spot/", controllers.CreateSpot).Methods("POST")
	router.HandleFunc("/room/{roomId:[0-9]+}/spot/", controllers.GetSpots).Methods("GET")

	router.HandleFunc("/spot/{spotId:[0-9]+}", controllers.GetSpot).Methods("GET")
	router.HandleFunc("/spot/{spotId:[0-9]+}", controllers.UpdateSpot).Methods("PUT")
	router.HandleFunc("/spot/{spotId:[0-9]+}", controllers.DeleteSpot).Methods("DELETE")
}
