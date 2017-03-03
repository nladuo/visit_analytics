package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

var (
	visit_chan chan Visit = make(chan Visit, 1) //no bufferred channel
)

func main() {

	InitDB()

	router := gin.Default()
	router.GET("/analytics.js", func(c *gin.Context) {
		go analyse(c)
		c.Header("Content-Type", "application/javascript")
		c.String(http.StatusOK, "console.log(\"https://github.com/nladuo/visit_analytics\")")
	})

	router.StaticFile("/test", "./www/test.html")

	go func() {
		for {
			visit := <-visit_chan
			HandleVisit(visit)
		}
	}()

	router.Run(":3000")
}

// 根据referer记录到数据库
func analyse(c *gin.Context) {
	if len(c.Request.Referer()) == 0 {
		return
	}

	title := getTitle(c.Request.Referer())
	fmt.Println(title)

	visit_chan <- Visit{
		ClientIp:  c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		Referer:   c.Request.Referer(),
		Title:     title,
	}
}

// 获取文章标题
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
