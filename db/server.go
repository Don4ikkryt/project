package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

var informationAboutUsers *mgo.Collection
var informationAboutAdmins *mgo.Collection

func serverDB() {

	router := gin.Default()
	router.POST("/updateUser", updateUserDB)
	router.POST("/createUser", createNewUserDB)
	router.GET("/deleteUser", deleteUserDB)
	router.GET("/readUser", readUser)
	router.POST("/autorizationUser", autorizationUser)
	router.POST("/actionAutorizationUser", actionAutorizationUser)

	router.POST("/createCat", createCat)
	router.POST("/updateCat", updateCat)
	router.GET("/deleteCat", deleteCat)

	router.POST("/message", writeMessage)
	router.GET("/readUsers", readUsers)
	router.GET("/readAdmin", readAdmin)
	router.POST("/actionAutorizationAdmin", actionAutorizationAdmin)
	router.POST("/createAdmin", createNewAdminDB)
	router.POST("/autorizationAdmin", autorizationAdmin)
	router.Run("localhost:3030")

}

func autorizationUser(c *gin.Context) {
	var u User
	err := c.Bind(&u)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	var user User
	if err := informationAboutUsers.Find(obj{"username": strings.ToLower(u.Username)}).One(&user); err == nil {
		if user.Password == u.Password {
			user.Username = encryptCookie(user.Username)
			user.Password = encryptCookie(user.Password)
			c.JSON(http.StatusOK, user)
			return
		}
		c.String(http.StatusBadRequest, "wrong password or username")
		return

	}

}
func actionAutorizationUser(c *gin.Context) {
	obj1 := obj{}
	err := c.Bind(&obj1)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	if value, ok := obj1["cookie"]; ok {
		muxForUserCookie.Lock()
		if _, ok := userCookie[value.(string)]; ok {
			muxForUserCookie.Unlock()
			c.String(http.StatusOK, "autorizated")
			return
		}
	}
	c.String(http.StatusBadRequest, "not autorizated")
}
func autorizationAdmin(c *gin.Context) {
	var a Admin
	err := c.Bind(&a)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	var admin Admin
	if err := informationAboutAdmins.Find(obj{"username": strings.ToLower(a.Username)}).One(&admin); err == nil {
		if admin.Password == a.Password {
			admin.Username = encryptCookie(admin.Username)
			admin.Password = encryptCookie(admin.Password)
			c.JSON(http.StatusOK, admin)
			return
		}
		c.String(http.StatusBadRequest, "wrong password or username")
		return

	}

}
func actionAutorizationAdmin(c *gin.Context) {
	obj1 := obj{}
	err := c.Bind(&obj1)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	if value, ok := obj1["cookie"]; ok {
		muxForAdminCookie.Lock()
		if _, ok := adminCookie[value.(string)]; ok {
			muxForAdminCookie.Unlock()
			c.String(http.StatusOK, "autorizated")
			return
		}
	}
	c.String(http.StatusBadRequest, "not autorizated")
}
