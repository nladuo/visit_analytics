package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ADMIN_UNAME = "nladuo" // the admin username
	ADMIN_PASS  = "nladuo" // the admin password
)

func MakeRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("frontend/templates/*")

	router.Static("/static/js", "./www/js")
	router.Static("/static/css", "./www/css")
	router.Static("/static/imgs", "./www/imgs")

	authorized := router.Group("/manage", gin.BasicAuth(gin.Accounts{
		ADMIN_UNAME: ADMIN_PASS,
	}))

	authorized.GET("/", manageTemplate)

	authorized.GET("/api/hosts", apiGetHosts)
	authorized.GET("/api/pages", apiGetPages)
	authorized.GET("/api/daily_records/", apiGetDailyRecords)
}

func manageTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "manage.tmpl", gin.H{
		"analytics_url": "http://localhost:3000/analytics.js",
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

func apiGetDailyRecords(c *gin.Context) {

}

func showResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
