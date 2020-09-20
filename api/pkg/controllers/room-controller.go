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

var NewRoom models.Room

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	CreateRoom := &models.Room{}
	utils.ParseBody(r, CreateRoom)
	b:= CreateRoom.CreateRoom()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	newRooms:= models.GetAllRooms()
	res, _ := json.Marshal(newRooms)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRoomById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	RoomId := vars["roomId"]
	ID, err:= strconv.ParseInt(RoomId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	RoomDetails, _:= models.GetRoomById(ID)
	res, _ := json.Marshal(RoomDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	var updateRoom = &models.Room{}
	utils.ParseBody(r, updateRoom)
	vars := mux.Vars(r)
	RoomId := vars["roomId"]
	ID, err:= strconv.ParseInt(RoomId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	RoomDetails, db:= models.GetRoomById(ID)
	if updateRoom.Name != "" {
		RoomDetails.Name = updateRoom.Name
	}
	db.Save(&RoomDetails)
	res, _ := json.Marshal(RoomDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	RoomId := vars["roomId"]
	ID, err:= strconv.ParseInt(RoomId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	// TO-Do: check if spots exists
	models.DeleteAllSpotByRoomId(ID)
	Room:= models.DeleteRoom(ID)
	res, _ := json.Marshal(Room)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
