package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name" form:"name" bson:"name"`
	Age      int    `json:"age" form:"age" bson:"age"`
	Gender   string `json:"gender" form:"gender" bson:"gender"`
	ID       uint   `json:"id" form:"id" bson:"_id"`
	Cats     []Cat  `json:"cats" form:"cats" bson:"cats"`
	Username string `json:"username" form:"username" bson:"username"`
	Password string `json:"password" form:"password" bson:"password"`
	Message  string `json:"message" form:"message" bson:"message"`
}

func readUser(c *gin.Context) {
	str := c.Query("id")
	if idOfReading, err := strconv.Atoi(str); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		if idOfReading > 0 {
			muxForUserCache.Lock()
			if user, ok := userCache[uint(idOfReading)]; ok {
				muxForUserCache.Unlock()
				user.Username = encryptCookie(user.Username)
				user.Password = encryptCookie(user.Password)
				c.JSON(http.StatusOK, user)
				return
			}
			c.String(http.StatusBadRequest, "not found")
			return
		}
	}

}

func createNewUserDB(c *gin.Context) {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		user.Message = ""
		user.ID = idUser + 1
		user.Username = strings.ToLower(user.Username)
		err := informationAboutUsers.Insert(user)
		if err != nil {
			c.String(http.StatusBadRequest, "this username already exists")
			return
		} else {
			idUser++
			iteratorIDUser()
			muxForUserCache.Lock()
			userCache[idUser] = user
			muxForUserCache.Unlock()
			muxForUserCookie.Lock()
			userCookie[encryptCookie(user.Password)] = [2]string{user.Username, user.Password}
			muxForUserCookie.Unlock()
			user.Username = encryptCookie(user.Username)
			user.Password = encryptCookie(user.Password)
			c.JSON(http.StatusOK, user)
			return
		}
	}

}

func iteratorIDUser() {

	informationAboutUsers.UpdateId(0, obj{"_id": 0, "ID": idUser})

}
func updateUserDB(c *gin.Context) {

	user := User{}
	err := c.Bind(&user)
	updatingID := user.ID
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		if updatingID > 0 {
			muxForUserCache.Lock()
			user.Username = userCache[updatingID].Username
			user.Password = userCache[updatingID].Password
			user.Message = userCache[updatingID].Message
			user.Cats = userCache[updatingID].Cats
			muxForUserCache.Unlock()
			err := informationAboutUsers.UpdateId(updatingID, user)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
			} else {
				muxForUserCache.Lock()
				userCache[updatingID] = user
				muxForUserCache.Unlock()
				user.Username = encryptCookie(user.Username)
				user.Password = encryptCookie(user.Password)
				c.JSON(http.StatusOK, user)
			}
		}
	}
}
func deleteUserDB(c *gin.Context) {
	str := c.Query("id")
	if idOfDeleting, err := strconv.Atoi(str); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		if idOfDeleting > 0 {
			err := informationAboutUsers.RemoveId(idOfDeleting)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
			} else {
				muxForUserCache.Lock()
				delete(userCache, uint(idOfDeleting))
				muxForUserCache.Unlock()
				c.String(http.StatusOK, "Deleted")
			}
		}
	}
}
