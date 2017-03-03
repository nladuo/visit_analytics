package main

import "time"

func HandleVisit(visit Visit) {
	recordHost(visit)
	recordPage(visit)
	recordDailyRecord(visit)
}

func recordHost(visit Visit) {
	var count int
	db := GetDB()
	db.Model(&Host{}).Where("host_name = ?", visit.Host).Count(&count)
	if count == 0 {
		db.Create(&Host{HostName: visit.Host})
	}
}

func recordPage(visit Visit) {
	var page Page
	db.Where("url = ?", visit.Referer).Find(&page)
	if page.TotalCount == 0 {
		db.Create(&Page{
			Host:       visit.Host,
			Url:        visit.Referer,
			Title:      visit.Title,
			TotalCount: 1,
		})
	} else { // count plus one
		page.TotalCount += 1
		db.Save(&page)
	}
}

func recordDailyRecord(visit Visit) {
	var daily_record DailyRecord
	today := time.Now().Format("2006-01-02")

	db.Where("url = ? && date = ?", visit.Referer, today).Find(&daily_record)
	if daily_record.Id == 0 {
		db.Create(&DailyRecord{
			Url:   visit.Referer,
			Date:  today,
			Count: 1,
		})
	} else { // count plus one
		daily_record.Count += 1
		db.Save(&daily_record)
	}
}
