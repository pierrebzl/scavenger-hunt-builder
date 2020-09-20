package controllers

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

	"github.com/pierrebzl/scavenger-hunt/pkg/models"
	"github.com/pierrebzl/scavenger-hunt/pkg/utils"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	CreateRoom := &models.Room{}
	utils.ParseBody(r, CreateRoom)
	// TO-DO - Error handling
	Room := CreateRoom.CreateRoom()
	utils.Respond(w, http.StatusOK, Room)
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	Rooms := models.GetAllRooms()
	utils.Respond(w, http.StatusOK, Rooms)
}

func GetRoomById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	RoomId := vars["roomId"]
	ID, _ := strconv.ParseInt(RoomId, 0, 0)
	Room, _ := models.GetRoomById(ID)
	// TO-DO - Error handling
	utils.Respond(w, http.StatusOK, Room)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	var updateRoom = &models.Room{}
	utils.ParseBody(r, updateRoom)
	vars := mux.Vars(r)
	RoomId := vars["roomId"]
	ID, _ := strconv.ParseInt(RoomId, 0, 0)
	RoomDetails, db := models.GetRoomById(ID)
	if updateRoom.Name != "" {
		RoomDetails.Name = updateRoom.Name
	}
	// TO-DO - Add fields to update
	db.Save(&RoomDetails)
	utils.Respond(w, http.StatusOK, RoomDetails)
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	RoomId := vars["roomId"]
	ID, _ := strconv.ParseInt(RoomId, 0, 0)
	// TO-DO: check if spots exists
	models.DeleteAllSpotByRoomId(ID)
	Room:= models.DeleteRoom(ID)
	utils.Respond(w, http.StatusOK, Room)
}
