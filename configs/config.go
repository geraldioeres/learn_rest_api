package configs

import (
	"learn_api/models/twitters"
	"learn_api/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:diosql@tcp(127.0.0.1:3306)/learn_api?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB failed to connect")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&users.User{}, &twitters.Tweet{})
}
