package models

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name        string `json:"name"`
	TeamRefer 	uint64
}
