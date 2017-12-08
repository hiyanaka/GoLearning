package model

import (
	"github.com/jinzhu/gorm"
)

var db = GetDB()

type User struct {
	Name     string
	Mail     string
	PassWord string
}

func GetDB() *gorm.DB {
	con, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	return con
}
