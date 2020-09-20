package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

	"github.com/pierrebzl/scavenger-hunt/pkg/models"
	"github.com/pierrebzl/scavenger-hunt/pkg/utils"
)

var NewSpot models.Spot

func GetSpots(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	RoomId := vars["roomId"]
	ID, err:= strconv.ParseInt(RoomId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	// To-DO Handle if room does not exist
	Spots := models.GetAllSpotsByRoomId(ID)
	res, _ := json.Marshal(Spots)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	SpotId := vars["spotId"]
	ID, err:= strconv.ParseInt(SpotId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	RoomDetails, _:= models.GetSpotById(ID)
	res, _ := json.Marshal(RoomDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	RoomId, err:= strconv.ParseUint(vars["roomId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	CreateSpot := &models.Spot{RoomRefer: RoomId}
	utils.ParseBody(r, CreateSpot)
	b:= CreateSpot.CreateSpot()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateSpot(w http.ResponseWriter, r *http.Request) {
	var updateSpot = &models.Spot{}
	utils.ParseBody(r, updateSpot)
	vars := mux.Vars(r)
	// TO-DO handle when spot id does not reexist
	SpotId := vars["spotId"]
	ID, err:= strconv.ParseInt(SpotId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	Spot, db:= models.GetSpotById(ID)
	if updateSpot.Name != "" {
		Spot.Name = updateSpot.Name
	}
	if updateSpot.Clue != "" {
		Spot.Clue = updateSpot.Clue
	}
	db.Save(&Spot)
	res, _ := json.Marshal(Spot)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func DeleteSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	SpotId := vars["spotId"]
	ID, err:= strconv.ParseInt(SpotId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	Spot:= models.DeleteSpot(ID)
	res, _ := json.Marshal(Spot)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
