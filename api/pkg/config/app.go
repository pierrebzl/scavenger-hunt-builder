package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db * gorm.DB
)

func Connect() {
	dsn := "pierre:dbopenup@tcp(127.0.0.1)/scavenger?charset=utf8&parseTime=True&loc=Local"
  	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
