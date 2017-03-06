package main

// table object mapping: hosts
type Host struct {
	Id       int    `gorm:"primary_key";"AUTO_INCREMENT"`
	HostName string `gorm:"unique"`
}

// table object mapping: pages
type Page struct {
	Id         int `gorm:"primary_key";"AUTO_INCREMENT"`
	Host       string
	Url        string `gorm:"unique"`
	Title      string
	TotalCount int
}

// table object mapping: daily_records
type DailyRecord struct {
	Id        int `gorm:"primary_key";"AUTO_INCREMENT"`
	Url       string
	Date      string
	TimeStamp int64
	Count     int
}

// table object mapping: monthly_record
type MonthlyRecord struct {
	Id        int `gorm:"primary_key";"AUTO_INCREMENT"`
	Url       string
	Date      string
	TimeStamp int64
	Count     int
}

type Visit struct {
	ClientIp  string
	UserAgent string
	Referer   string
	Title     string
	Host      string
}
