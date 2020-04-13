package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Cat struct {
	Name   string `json:"name" form:"name" bson:"name"`
	Age    int    `json:"age" form:"age" bson:"age"`
	Breed  string `json:"breed" form:"breed" bson:"breed"`
	Gender string `json:"gender" form:"gender" bson:"gender"`
	Color  string `json:"color" form:"color" bson:"color"`
	ID     uint   `json:"id" form:"id" bson:"_id"`
}

func createCat(c *gin.Context) {
	cat := Cat{}
	c.Bind(&cat)
	autorizationData, err := c.Cookie("autorizationData")
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", err.Error())
		return
	}
	idStr, err := c.Cookie("id")
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", err.Error())
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", err.Error())
		return
	}
	user := User{}
	user.ID = uint(id)
	status, msg := requestActionAutorizationUser(autorizationData)
	if status == http.StatusOK {
		status, msg, userResp := requestCatCreating(&cat, int(user.ID))
		if status != http.StatusOK {
			c.HTML(http.StatusBadRequest, "message.html", &User{Name: msg})
			return
		}
		c.HTML(http.StatusOK, "userPage.html", userResp)
	} else {
		c.HTML(http.StatusOK, "message.html", &User{Name: msg})
	}
}
func updateCat(c *gin.Context) {
	cat := Cat{}
	c.Bind(&cat)
	autorizationData, err := c.Cookie("autorizationData")
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", err.Error())
		return
	}
	idStr, err := c.Cookie("id")
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", err.Error())
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", err.Error())
		return
	}
	user := User{}
	user.ID = uint(id)
	status, msg := requestActionAutorizationUser(autorizationData)
	if status == http.StatusOK {
		status, msg, userResp := requestCatUpdating(&cat, int(user.ID))
		if status != http.StatusOK {
			c.HTML(http.StatusBadRequest, "message.html", &User{Name: msg})
			return
		}
		c.HTML(http.StatusOK, "userPage.html", userResp)
	} else {
		c.HTML(http.StatusOK, "message.html", &User{Name: msg})
	}
}
func deleteCat(c *gin.Context) {
	str := c.Query("id")
	if idOfCat, err := strconv.Atoi(str); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		autorizationData, err := c.Cookie("autorizationData")
		if err != nil {
			c.HTML(http.StatusBadRequest, "message.html", err.Error())
			return
		}
		idStr, err := c.Cookie("id")
		if err != nil {
			c.HTML(http.StatusBadRequest, "message.html", err.Error())
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.HTML(http.StatusBadRequest, "message.html", err.Error())
			return
		}
		user := User{}
		user.ID = uint(id)
		status, msg := requestActionAutorizationUser(autorizationData)
		if status == http.StatusOK {
			status, msg, userResp := requestCatDeleting(idOfCat, int(user.ID))
			if status != http.StatusOK {
				c.HTML(http.StatusBadRequest, "message.html", &User{Name: msg})
				return
			}
			c.HTML(http.StatusOK, "userPage.html", userResp)
		} else {
			c.HTML(http.StatusOK, "message.html", &User{Name: msg})
		}
	}
}
