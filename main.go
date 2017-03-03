package main

import (
	"net/http"

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
	router.StaticFile("/test2", "./www/test.html")

	go func() { //handle visit record
		for {
			visit := <-visit_chan
			HandleVisit(visit)
		}
	}()

	router.Run(":3000")
}

// record according to Request.Referer()
func analyse(c *gin.Context) {
	referer := c.Request.Referer()
	if len(referer) == 0 {
		return
	}

	title := GetTitle(referer)
	host_name := GetHostName(referer)
	if host_name == "" {
		return
	}

	visit_chan <- Visit{
		ClientIp:  c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		Referer:   referer,
		Title:     title,
		Host:      host_name,
	}
}
