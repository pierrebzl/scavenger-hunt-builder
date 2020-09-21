package models

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name        string 	 `json:"name"`
	Color		string 	 `json:"color"`
	Players 	[]Player `gorm:"foreignKey:TeamRefer"`
}
