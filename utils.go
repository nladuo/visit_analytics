package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func TrimUrl(url string) string {
	url = strings.TrimRight(url, "/")
	url = strings.TrimRight(url, "/?")
	return url
}

// parse hostname from url
func GetHostName(url string) string {
	host_name := strings.Split(url, "//")[1]

	strs := strings.Split(host_name, "/")
	if len(strs) == 0 {
		return ""
	}
	return strs[0]
}

// get title from database or by httpClient
func GetTitle(url string) string {
	// get title from database
	page := findPage(url)
	if page.Id != 0 {
		return page.Title
	}

	// get title by httpClient
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return url
	}
	title := doc.Find("title").Text()
	title = strings.Trim(title, " ")
	if len(title) == 0 {
		return url
	}
	return title
}
