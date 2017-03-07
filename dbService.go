package main

import (
	"math/rand"
	"strconv"
	"time"
)

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
			Url:       visit.Referer,
			Date:      today,
			TimeStamp: time.Now().Unix(),
			Count:     1,
		})
	} else { // count plus one
		daily_record.Count += 1
		db.Save(&daily_record)
	}
}

func recordMonthlyRecord(visit Visit) {
	var monthly_record MonthlyRecord
	this_month := time.Now().Format("2006-01")

	db.Where("url = ? && date = ?", visit.Referer, this_month).Find(&monthly_record)
	if monthly_record.Id == 0 {
		db.Create(&MonthlyRecord{
			Url:       visit.Referer,
			Date:      this_month,
			TimeStamp: time.Now().Unix(),
			Count:     1,
		})
	} else { // count plus one
		monthly_record.Count += 1
		db.Save(&monthly_record)
	}
}

func findPage(url string) Page {
	var page Page
	db := GetDB()
	db.Where("url = ?", url).Find(&page)

	return page
}

func findHosts() []Host {
	hosts := []Host{}

	db := GetDB()
	db.Find(&hosts)
	return hosts
}

func findPages(host string) []Page {
	pages := []Page{}

	db := GetDB()
	db.Where("host = ?", host).Find(&pages)

	return pages
}

func findAllPages() []Page {
	pages := []Page{}

	db := GetDB()
	db.Find(&pages)

	return pages
}

func searchDailyRecords(url string, tm time.Time) []DailyRecord {
	dayly_records := []DailyRecord{}

	tomorrow := tm.AddDate(0, 0, 1)
	last_month := tm.AddDate(0, -1, -1)

	db := GetDB()
	db.Order("time_stamp asc").Where("url = ? && time_stamp <= ? && time_stamp >= ?",
		url, tomorrow.Unix(), last_month.Unix()).Find(&dayly_records)

	return dayly_records
}

func searchMonthlyRecords(url string, tm time.Time) []MonthlyRecord {
	monthly_records := []MonthlyRecord{}

	tomorrow := tm.AddDate(0, 0, 1)
	last_year := tm.AddDate(-1, -1, 0)

	db := GetDB()

	db.Order("time_stamp asc").Where("url = ? && time_stamp <= ? && time_stamp >= ?",
		url, tomorrow.Unix(), last_year.Unix()).Find(&monthly_records)

	return monthly_records
}

func generateRandomRecords() {
	now_str := strconv.FormatInt(time.Now().UnixNano(), 10)
	url := "http://localhost:3000/" + now_str
	db := GetDB()

	//create host
	db.Create(&Host{
		HostName: "http://localhost:3000/",
	})

	//create dailyrecord record
	now := time.Now()
	total_count := 0
	for i := 0; i < 40; i++ {
		tm := now.AddDate(0, 0, i*(-1))
		rand_num := rand.Intn(100)
		total_count += rand_num
		db.Create(&DailyRecord{
			Url:       url,
			Date:      tm.Format("2006-01-02"),
			TimeStamp: tm.Unix(),
			Count:     rand_num,
		})

	}

	//create page
	db.Create(&Page{
		Host:       "http://localhost:3000/",
		Url:        url,
		Title:      now_str,
		TotalCount: total_count,
	})

}
