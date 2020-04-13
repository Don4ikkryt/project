package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Cat struct {
	Name   string `json:"name" form:"name" bson:"name"`
	Age    int    `json:"age" form:"age" bson:"age"`
	Color  string `json:"color" form:"color" bson:"color"`
	Breed  string `json:"breed" form:"breed" bson:"breed"`
	Gender string `json:"gender" form:"gender" bson:"gender"`
	ID     uint   `json:"id" form:"id" bson:"_id"`
}

func createCat(c *gin.Context) {
	cat := Cat{}
	str := c.Query("id")
	if idOfUser, err := strconv.Atoi(str); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		err := c.Bind(&cat)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else {
			idOfCat := defineIdForCat(idOfUser)
			cat.ID = idOfCat
			err := informationAboutUsers.UpdateId(idOfUser, obj{"$push": obj{"cats": cat}})
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			muxForUserCache.Lock()
			if user, ok := userCache[uint(idOfUser)]; ok {
				user.Cats = append(user.Cats, cat)
				c.JSON(http.StatusOK, user)
			}
			muxForUserCache.Unlock()
		}
	}
}
func defineIdForCat(idOfUser int) (nextCatId uint) {
	muxForUserCache.Lock()
	if value, ok := userCache[uint(idOfUser)]; ok {
		for _, cat := range value.Cats {
			if cat.ID > nextCatId {
				nextCatId = cat.ID
				break
			}

		}
		muxForUserCache.Unlock()
		nextCatId++
		return nextCatId
	}
	return
}
func updateCat(c *gin.Context) {
	idUserStr := c.Query("id")
	if idOfUser, err := strconv.Atoi(idUserStr); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		cat := Cat{}
		err := c.Bind(&cat)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else {
			muxForUserCache.Lock()
			if user, ok := userCache[uint(idOfUser)]; ok {
				muxForUserCache.Unlock()
				i := 0
				for _, value := range user.Cats {
					if value.ID == cat.ID {
						err := informationAboutUsers.UpdateId(idOfUser, obj{"$pull": obj{"cats": obj{"_id": value.ID}}})
						if err != nil {

							c.String(http.StatusBadRequest, err.Error())
							return
						}
						err = informationAboutUsers.UpdateId(idOfUser, obj{"$push": obj{"cats": cat}})
						if err != nil {

							c.String(http.StatusBadRequest, err.Error())
							return
						}
						user.Cats[i] = cat
						c.JSON(http.StatusOK, user)
						return
					}
					i++
				}
				idCatStr := strconv.Itoa(int(cat.ID))
				c.String(http.StatusBadRequest, "no such cat with this id: "+idCatStr)

				return
			} else {
				c.String(http.StatusBadRequest, "no such user with this id")
				muxForUserCache.Unlock()
			}
		}

	}
}

func deleteCat(c *gin.Context) {
	idUserStr := c.Query("idOfUser")
	idOfUser, err := strconv.Atoi(idUserStr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	idCatStr := c.Query("idOfCat")
	idOfCat, err := strconv.Atoi(idCatStr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	muxForUserCache.Lock()
	if user, ok := userCache[uint(idOfUser)]; ok {
		muxForUserCache.Unlock()
		i := 0
		for _, value := range user.Cats {
			if int(value.ID) == idOfCat {
				err := informationAboutUsers.UpdateId(idOfUser, obj{"$pull": obj{"cats": obj{"_id": value.ID}}})
				if err != nil {

					c.String(http.StatusBadRequest, err.Error())
					return
				}
				user.Cats = append(user.Cats[:i], user.Cats[i+1:]...)

				c.JSON(http.StatusOK, user)
				return
			}
			i++
		}

		c.String(http.StatusBadRequest, "no such cat with this id: "+idCatStr)
		return
	} else {
		muxForUserCache.Unlock()
		c.String(http.StatusBadRequest, "no such user with this id")
	}
}
