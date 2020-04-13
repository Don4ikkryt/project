package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func server() {
	router := gin.Default()
	router.POST("/message", writeMessage)
	router.GET("/readUsers", readUsers)
	router.POST("/autorization", autorization)
	router.POST("/createAdmin", createAdmin)
	router.GET("/return", returnToAdminPage)
	router.LoadHTMLFiles("registrationForm.html", "adminPage.html", "retryRegistration.html", "message.html", "users.html")
	router.Run("localhost:4040")
}

func autorization(c *gin.Context) {
	admin := Admin{}
	err := c.Bind(&admin)
	if err != nil {
		c.HTML(http.StatusBadRequest, "retryRegistration.html", &Admin{Name: err.Error()})
		return
	}

	status, msg, adminResp := requestAutorization(&admin)
	if status == http.StatusOK {
		c.SetCookie("id", strconv.Itoa(int(adminResp.ID)), 0, "", "", http.SameSiteDefaultMode, false, true)
		c.SetCookie("autorizationData", adminResp.Password, 0, "", "", http.SameSiteDefaultMode, false, true)
		c.HTML(status, "adminPage.html", adminResp)
		return
	}

	c.HTML(http.StatusBadRequest, "retryRegistration.html", &Admin{Name: msg})
	return
}

func returnToAdminPage(c *gin.Context) {
	autorizationData, err := c.Cookie("autorizationData")
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", &Admin{Name: err.Error()})
		return
	}
	idStr, err := c.Cookie("id")
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", &Admin{Name: err.Error()})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", &Admin{Name: err.Error()})
		return
	}
	_, _, adminResp := requestReading(id)
	if adminResp != nil {
		status, msg := requestActionAutorizationAdmin(autorizationData)
		if status == http.StatusOK {
			c.HTML(http.StatusOK, "adminPage.html", &adminResp)
		} else {
			c.HTML(status, "retryRegistration.html", &Admin{Name: msg})
		}
	} else {
		c.HTML(http.StatusBadRequest, "retryRegistration.html", &Admin{Name: "something go wrong"})
	}
}
