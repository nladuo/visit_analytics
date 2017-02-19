package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	router := gin.Default()
	router.GET("/analytics.js", func(c *gin.Context) {
		analyse(c)
		c.Header("Content-Type", "application/javascript")
		c.String(http.StatusOK, "console.log(\"https://github.com/nladuo/visit_analytics\")")
	})

	router.StaticFile("/index", "./www/index.html")

	router.Run(":3000")
}

func analyse(c *gin.Context) {
	if len(c.Request.Referer()) == 0 {
		return
	}
	fmt.Println("ClientIp:", c.ClientIP())
	fmt.Println("User-Agent:", c.Request.UserAgent())
	fmt.Println("Referer:", c.Request.Referer())
	fmt.Println("Title:", getTitle(c.Request.Referer()))
}

func getTitle(url string) string {
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
