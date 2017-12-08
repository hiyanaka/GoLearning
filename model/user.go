package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db = GetDB()

type User struct {
	gorm.Model
	Name     string
	Mail     string
	PassWord string
}

func GetDB() *gorm.DB {
	con, err := gorm.Open("mysql", "root@(localhost:3306)/nagaoka?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return con
}

func CreateUser(name, mail, password string) {
	db.Create(&User{
		Name:     name,
		Mail:     mail,
		PassWord: password,
	})

}

func LoginCheck(mail, password string) bool {
	var users User
	db.Find(&users)
	db.Where("mail = ?", mail).First(&users)
	if users.Mail != mail {
		fmt.Println("mail error")
		return false
	} else if users.PassWord != password {
		fmt.Println("password error")
		return false
	} else {
		fmt.Println("OK")
		return true
	}

	// fmt.Println(users.PassWord)
}

func DeleteUser() {

}

func UserMigrate() {
	db.AutoMigrate(&User{})
}
