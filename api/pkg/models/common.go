package models

import (
	"github.com/pierrebzl/scavenger-hunt/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Room{})
	db.AutoMigrate(&Spot{})
	db.AutoMigrate(&Player{})
	db.AutoMigrate(&Team{})
}
