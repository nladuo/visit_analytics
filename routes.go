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

	router.GET("/", loginTemplate)
	router.GET("/manage", manageTemplate)

	router.POST("/api/login", apiLogin)
	router.GET("/api/hosts", apiGetHosts)
	router.GET("/api/pages", apiGetPages)
	router.GET("/api/daily_records", apiGetDailyRecords)

}

func loginTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"analytics_url": "http://localhost:3000/analytics.js",
	})
}

func manageTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "manage.tmpl", gin.H{
		"analytics_url": "http://localhost:3000/analytics.js",
	})
}

func apiLogin(c *gin.Context) {

}

func apiGetHosts(c *gin.Context) {

}

func apiGetPages(c *gin.Context) {

}

func apiGetDailyRecords(c *gin.Context) {

}
