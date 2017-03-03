package main

import (
	"fmt"
	"time"
)

func HandleVisit(visit Visit) {
	addHost(visit)
	addDailyRecord(visit)
	addPage(visit)
}

func addHost(visit Visit) {
	host_name := GetHostName(visit.Referer)
	db := GetDB()
	var count int
	db.Model(&Host{}).Where("host_name = ?", host_name).Count(&count)
	if count == 0 {
		db.Create(&Host{HostName: host_name})
	}
}

func addPage(visit Visit) {

}

func addDailyRecord(visit Visit) {
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
}
