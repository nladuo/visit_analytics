package main

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDB() {
	err := connectDB()
	if err != nil {
		panic(err)
	}

	if (!db.HasTable(&Host{})) {
		db.CreateTable(&Host{})
	}

	if (!db.HasTable(&Page{})) {
		db.CreateTable(&Page{})
	}

	if (!db.HasTable(&DailyRecord{})) {
		db.CreateTable(&DailyRecord{})
	}

	if (!db.HasTable(&MonthlyRecord{})) {
		db.CreateTable(&MonthlyRecord{})
	}

}

func connectDB() error {
	config := GetConfig()
	var err error
	db, err = gorm.Open("mysql", config.DB.Username+":"+config.DB.Password+
		"@tcp("+config.DB.Host+":"+config.DB.Host+")/"+config.DB.DBName+"?charset=utf8&parseTime=True")

	return err
}

func GetDB() *gorm.DB {
	if db == nil {
		err := connectDB()
		if err != nil {
			panic(err)
		}
	}
	return db
}
