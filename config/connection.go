package config

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres",
		"host=localhost port=5432 user=postgres password=postgres dbname=todogolang sslmode=disable")

	if err != nil {
		panic(err.Error())
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
