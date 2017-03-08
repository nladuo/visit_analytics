package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MakeRoutes(router *gin.Engine) {
	config := GetConfig()

	router.LoadHTMLGlob("frontend/templates/*")

	router.StaticFS("/static", http.Dir("./www"))

	authorized := router.Group("/manage", gin.BasicAuth(gin.Accounts{
		config.Manage.Username: config.Manage.Password,
	}))

	authorized.GET("/", manageTemplate)

	authorized.GET("/api/hosts", apiGetHosts)
	authorized.GET("/api/pages", apiGetPages)
	authorized.GET("/api/records", apiGetRecords)
}

func manageTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "manage.tmpl", gin.H{
		"analytics_url": "http://" + config.DeployHost + "/analytics.js",
	})
}

func apiGetHosts(c *gin.Context) {
	showResponse(c, 0, "success", findHosts())
}

func apiGetPages(c *gin.Context) {
	host := c.DefaultQuery("host", "")
	if host == "all" {
		showResponse(c, 0, "success", findAllPages())
	} else {
		showResponse(c, 0, "success", findPages(host))
	}
}

func apiGetRecords(c *gin.Context) {
	url := c.DefaultQuery("url", "")
	_type := c.DefaultQuery("type", "0")
	date := c.DefaultQuery("date", "")

	//parse time
	tm, err := time.Parse("2006-01-02", date)
	if err != nil {
		showResponse(c, 1, "parse date error", "")
		fmt.Println(err.Error())
		return
	}

	if _type == "0" {
		showResponse(c, 0, "success", searchDailyRecords(url, tm))
	} else if _type == "1" {
		showResponse(c, 0, "success", searchMonthlyRecords(url, tm))
	} else {
		showResponse(c, 1, "error paramers", "")
	}
}

func showResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
