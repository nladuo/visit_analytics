package main

type Host struct {
	Id       int    `gorm:"primary_key";"AUTO_INCREMENT"`
	HostName string `gorm:"unique"`
}

type Page struct {
	Id         int `gorm:"primary_key";"AUTO_INCREMENT"`
	Host       string
	Url        string `gorm:"unique"`
	Title      string
	TotalCount int
}

type DailyRecord struct {
	Id    int `gorm:"primary_key";"AUTO_INCREMENT"`
	Url   int
	Date  string
	Count int
}

type Visit struct {
	ClientIp  string
	UserAgent string
	Referer   string
	Title     string
}
