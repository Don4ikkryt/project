package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Admin struct {
	Name     string `json:"name" form:"name" bson:"name"`
	ID       uint   `json:"id" form:"id" bson:"_id"`
	Username string `json:"username" form:"username" bson:"username"`
	Password string `json:"password" form:"password" bson:"password"`
}

func writeMessage(c *gin.Context) {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	err = informationAboutUsers.UpdateId(user.ID, obj{"$set": obj{"message": user.Message}})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		muxForUserCache.Lock()
		if value, ok := userCache[user.ID]; ok {
			value.Message = user.Message
			userCache[user.ID] = value
		}
		muxForUserCache.Unlock()
		c.String(http.StatusOK, "writed")
	}

}
func readUsers(c *gin.Context) {
	users := make([]User, 0)
	muxForUserCache.Lock()
	for _, value := range userCache {
		value.Password = ""
		users = append(users, value)
	}
	muxForUserCache.Unlock()
	c.JSON(http.StatusOK, users)
}
func createNewAdminDB(c *gin.Context) {
	admin := Admin{}
	err := c.Bind(&admin)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		admin.ID = idAdmin + 1
		admin.Username = strings.ToLower(admin.Username)
		err := informationAboutAdmins.Insert(admin)
		if err != nil {
			c.String(http.StatusBadRequest, "this username already exists")
			return
		} else {
			idAdmin++
			iteratorIDAdmin()
			muxForAdminCache.Lock()
			adminCache[idAdmin] = admin
			muxForAdminCache.Unlock()
			muxForAdminCookie.Lock()
			adminCookie[encryptCookie(admin.Password)] = [2]string{admin.Username, admin.Password}
			muxForAdminCookie.Unlock()
			admin.Username = encryptCookie(admin.Username)
			admin.Password = encryptCookie(admin.Password)
			c.JSON(http.StatusOK, admin)
			return
		}
	}

}
func iteratorIDAdmin() {

	informationAboutAdmins.UpdateId(0, obj{"_id": 0, "ID": idAdmin})

}

func readAdmin(c *gin.Context) {
	str := c.Query("id")
	if idOfReading, err := strconv.Atoi(str); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		if idOfReading > 0 {
			muxForAdminCache.Lock()
			if admin, ok := adminCache[uint(idOfReading)]; ok {
				muxForAdminCache.Unlock()
				admin.Username = encryptCookie(admin.Username)
				admin.Password = encryptCookie(admin.Password)
				c.JSON(http.StatusOK, admin)
				return
			}
			c.String(http.StatusBadRequest, "not found")
			return
		}
	}

}
