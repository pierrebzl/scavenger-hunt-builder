package models

import (
	"github.com/pierrebzl/scavenger-hunt/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Room struct {
	gorm.Model
	Name        string `json:"name"`
	Spots 		[]Spot `gorm:"foreignKey:RoomRefer"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Room{})
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
	var getRoom Room
	db:=db.Where("Id = ?", Id).Find(&getRoom)
	return &getRoom, db
}

func DeleteRoom(Id int64) Room {
	var Room Room
	db.Where("Id = ?", Id).Delete(Room)
	return Room
}
