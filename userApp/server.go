package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func server() {

	router := gin.Default()
	router.POST("/updateUser", updateUser)
	router.POST("/createUser", createUser)
	router.GET("/deleteUser", deleteUser)
	router.POST("/autorization", autorization)
	router.POST("/updateCat", updateCat)
	router.POST("/createCat", createCat)
	router.GET("/deleteCat", deleteCat)
	router.GET("/return", returnToUserPage)
	router.LoadHTMLFiles("registrationForm.html", "userPage.html", "retryRegistration.html", "message.html")
	router.Run()
}
func autorization(c *gin.Context) {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		c.HTML(http.StatusBadRequest, "retryRegistration.html", &User{Name: err.Error()})
		return
	}

	status, msg, userResp := requestAutorization(&user)
	if status == http.StatusOK {
		c.SetCookie("id", strconv.Itoa(int(userResp.ID)), 0, "", "", http.SameSiteDefaultMode, false, true)
		c.SetCookie("autorizationData", userResp.Password, 0, "", "", http.SameSiteDefaultMode, false, true)
		c.HTML(status, "userPage.html", userResp)
		return
	}

	c.HTML(http.StatusBadRequest, "retryRegistration.html", &User{Name: msg})
	return
}
func returnToUserPage(c *gin.Context) {
	autorizationData, err := c.Cookie("autorizationData")
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", &User{Name: err.Error()})
		return
	}
	idStr, err := c.Cookie("id")
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", &User{Name: err.Error()})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", &User{Name: err.Error()})
		return
	}
	_, _, userResp := requestReading(id)
	if userResp != nil {
		status, msg := requestActionAutorizationUser(autorizationData)
		if status == http.StatusOK {
			c.HTML(http.StatusOK, "userPage.html", &userResp)
		} else {
			c.HTML(status, "retryRegistration.html", &User{Name: msg})
		}
	} else {
		c.HTML(http.StatusBadRequest, "retryRegistration.html", &User{Name: "something go wrong"})
	}
}
