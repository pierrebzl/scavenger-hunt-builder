package models

import (
	"github.com/pierrebzl/scavenger-hunt/pkg/config"
	"gorm.io/gorm"
)

type Spot struct {
	gorm.Model
	Name 		string `json:"name"`
	Clue		string `json:"clue"`
	RoomRefer	uint64
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Spot{})
	var spots []Spot
	db.Model(&spots).Association("Spots")
}

func (b *Spot) CreateSpot() *Spot {
	db.Create(&b)
	return b
}

func GetAllSpotsByRoomId(Id int64) []Spot {
	var Spots []Spot
	db.Find(&Spots).Where("Id = ?", Id)
	return Spots
}

func GetSpotById(Id int64) (*Spot , *gorm.DB) {
	var Spot Spot
	db:=db.Where("Id = ?", Id).Find(&Spot)
	return &Spot, db
}

func DeleteSpot(Id int64) Spot {
	var Spot Spot
	db.Where("Id = ?", Id).Delete(Spot)
	return Spot
}

func DeleteAllSpotByRoomId(Id int64) Spot {
	var Spot Spot
	db.Where("room_refer = ?", Id).Delete(Spot)
	return Spot
}
