package controllers

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

	"github.com/pierrebzl/scavenger-hunt/pkg/models"
	"github.com/pierrebzl/scavenger-hunt/pkg/utils"
)

func GetSpots(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	RoomId := vars["roomId"]
	ID, _ := strconv.ParseInt(RoomId, 0, 0)
	// TO-DO Handle if room does not exist
	Spots := models.GetAllSpotsByRoomId(ID)
	utils.Respond(w, http.StatusOK, Spots)
}

func GetSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	SpotId := vars["spotId"]
	ID, _ := strconv.ParseInt(SpotId, 0, 0)
	Spot, _:= models.GetSpotById(ID)
	utils.Respond(w, http.StatusOK, Spot)
}

func CreateSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	RoomId, _ := strconv.ParseUint(vars["roomId"], 0, 0)
	// TO-DO - Err if room does not exist
	CreateSpot := &models.Spot{RoomRefer: RoomId}
	utils.ParseBody(r, CreateSpot)
	NewSpot := CreateSpot.CreateSpot()
	utils.Respond(w, http.StatusOK, NewSpot)
}

func UpdateSpot(w http.ResponseWriter, r *http.Request) {
	var updateSpot = &models.Spot{}
	utils.ParseBody(r, updateSpot)
	vars := mux.Vars(r)
	// TO-DO handle when spot id does not reexist
	SpotId := vars["spotId"]
	ID, _ := strconv.ParseInt(SpotId, 0, 0)
	Spot, db:= models.GetSpotById(ID)
	if updateSpot.Name != "" {
		Spot.Name = updateSpot.Name
	}
	if updateSpot.Clue != "" {
		Spot.Clue = updateSpot.Clue
	}
	db.Save(&Spot)
	utils.Respond(w, http.StatusOK, Spot)
}

func DeleteSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	SpotId := vars["spotId"]
	ID, _ := strconv.ParseInt(SpotId, 0, 0)
	Spot:= models.DeleteSpot(ID)
	utils.Respond(w, http.StatusOK, Spot)
}
