package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name        string `json:"name"`
	Spots 		[]Spot `gorm:"foreignKey:RoomRefer"`
}

func (r *Room) CreateRoom() *Room {
	db.Create(&r)
	return r
}

func GetAllRooms() []Room {
	var Rooms []Room
	db.Find(&Rooms)
	return Rooms
}

func GetRoomById(Id int64) (*Room , *gorm.DB){
	var Room Room
	db:=db.Where("Id = ?", Id).Find(&Room)
	return &Room, db
}

func DeleteRoom(Id int64) Room {
	var Room Room
	db.Where("Id = ?", Id).Delete(Room)
	return Room
}
