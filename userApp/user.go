package main

import (
	"net/http"
	"strconv"

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

func updateUser(c *gin.Context) {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", &User{Name: err.Error()})
		return
	}
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
	updatingID, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", &User{Name: err.Error()})
		return
	}

	if updatingID > 0 {
		status, msgRespAutorization := requestActionAutorizationUser(autorizationData)
		if status == http.StatusOK {
			user.ID = uint(updatingID)
			status, msgResp, userResp := requestUpdatingUserInDB(&user)
			if status == http.StatusOK {
				c.HTML(status, "userPage.html", &userResp)
				return
			}
			c.HTML(status, "message.html", &User{Name: msgResp})
			return
		} else {
			c.HTML(status, "retryRegistration.html", &User{Name: msgRespAutorization})
		}
	} else {
		c.HTML(http.StatusBadRequest, "message.html", &User{Name: "Negative id"})
	}
}

func createUser(c *gin.Context) {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		c.HTML(http.StatusBadRequest, "retryRegistraion.html", &User{Name: err.Error()})
	} else {
		status, msg, userResp := requestCreatingUserInDB(&user)
		if status != http.StatusOK {
			c.HTML(http.StatusBadRequest, "retryRegistration.html", &User{Name: msg})
			return
		}
		c.SetCookie("autorizationData", userResp.Password, 0, "", "", http.SameSiteDefaultMode, false, true)
		c.SetCookie("id", strconv.Itoa(int(userResp.ID)), 0, "", "", http.SameSiteDefaultMode, false, true)
		c.HTML(http.StatusOK, "userPage.html", userResp)

	}
}

func deleteUser(c *gin.Context) {
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
	idOfDeleting, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "message.html", &User{Name: err.Error()})
		return
	}
	status, msg := requestActionAutorizationUser(autorizationData)
	if status == http.StatusOK {
		if idOfDeleting > 0 {

			status, msg := requestDeletingUserInDB(idOfDeleting)
			if status != http.StatusOK {
				c.HTML(status, "message.html", &User{Name: msg})
			}

			c.HTML(http.StatusOK, "retryRegistration.html", &User{Name: msg})

		} else {
			c.HTML(http.StatusBadRequest, "message.html", &User{Name: "Negative id"})
		}
	} else {

		c.HTML(http.StatusOK, "retryRegistration.html", &User{Name: msg})
	}
}
