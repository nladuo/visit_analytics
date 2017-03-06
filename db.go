package main

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jinzhu/gorm"
)

const (
	DB_USER = "root"
	DB_PASS = "root"
	DB_HOST = "localhost"
	DB_PORT = "3306"
	DBNAME  = "visit_analytics"
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
	var err error
	db, err = gorm.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DBNAME+"?charset=utf8&parseTime=True")

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
