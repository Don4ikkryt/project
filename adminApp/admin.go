package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Users []User
}
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
		c.HTML(http.StatusBadRequest, "message.html", &Admin{Name: err.Error()})
		return
	}
	autorizationData, err := c.Cookie("autorizationData")
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", &User{Name: err.Error()})
		return
	}
	status, msg := requestActionAutorizationAdmin(autorizationData)
	if status == http.StatusOK {
		status, msg := requestWriteMessage(int(user.ID), user.Message)
		c.HTML(status, "message.html", &Admin{Name: msg})

	} else {
		c.HTML(status, "message.html", &Admin{Name: msg})
	}
}

func readUsers(c *gin.Context) {
	autorizationData, err := c.Cookie("autorizationData")
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", &Admin{Name: err.Error()})
		return
	}
	status, msg := requestActionAutorizationAdmin(autorizationData)
	if status == http.StatusOK {
		status, msg, users := requestReadingUsers()
		if status == http.StatusOK {
			c.HTML(http.StatusOK, "users.html", &Users{Users: users})
			return
		} else {
			c.HTML(status, "message.html", &Admin{Name: msg})
		}
	} else {
		c.HTML(status, "message.html", &Admin{Name: msg})
	}

}
func createAdmin(c *gin.Context) {
	admin := Admin{}
	err := c.Bind(&admin)
	if err != nil {
		c.HTML(http.StatusBadRequest, "retryRegistraion.html", &Admin{Name: err.Error()})
	} else {
		status, msg, adminResp := requestCreatingAdminInDB(&admin)
		if status != http.StatusOK {
			c.HTML(http.StatusBadRequest, "retryRegistration.html", &Admin{Name: msg})
			return
		}
		c.SetCookie("autorizationData", adminResp.Password, 0, "", "", http.SameSiteDefaultMode, false, true)
		c.SetCookie("id", strconv.Itoa(int(adminResp.ID)), 0, "", "", http.SameSiteDefaultMode, false, true)
		c.HTML(http.StatusOK, "adminPage.html", adminResp)

	}
}
