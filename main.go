package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Cat struct {
	Name   string `json:"name" form:"name"`
	Age    int    `json:"age" form:"age"`
	Color  string `json:"color" form:"color"`
	Breed  string `json:"breed" form:"breed"`
	Gender string `json:"gender" form:"gender"`
	ID     uint   `json:"id" form:"id"`
}

var catsDB map[uint]Cat
var id uint

func main() {
	router := gin.Default()
	catsDB = make(map[uint]Cat, 30)
	router.LoadHTMLFiles("catFormer.html", "catRead.html", "catCreate.html", "catUpdate.html", "catDelete.html")

	router.GET("/", hendlerFunc)
	router.POST("/update", updateCat)
	router.POST("/create", createNewCat)
	router.GET("/read", readCat)
	router.GET("/delete", deleteCat)

	router.Run()
}

func hendlerFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "catFormer.html", nil)
}
func updateCat(c *gin.Context) {

	cat := Cat{}
	err := c.Bind(&cat)
	updatingID := cat.ID
	if err != nil {
		c.HTML(http.StatusOK, "catUpdate.html", Cat{Name: "You entered wrong data"})
	} else {
		if updatingID > 0 {
			_, ok := catsDB[uint(updatingID)]
			if ok {
				catsDB[uint(updatingID)] = cat
				c.HTML(http.StatusOK, "catUpdate.html", cat)
			} else {
				c.HTML(http.StatusOK, "catUpdate.html", Cat{Name: "There is no such cat to be updated"})

			}
		}
	}

}
func createNewCat(c *gin.Context) {
	cat := Cat{}
	err := c.Bind(&cat)
	if err != nil {
		c.HTML(http.StatusOK, "catCreate.html", Cat{})
	} else {
		id++
		cat.ID = id
		catsDB[uint(id)] = cat
		c.HTML(http.StatusOK, "catCreate.html", cat)
	}

}
func deleteCat(c *gin.Context) {
	str := c.Query("id")
	if idOfDeleting, err := strconv.Atoi(str); err != nil {
		c.HTML(http.StatusOK, "catDelete.html", Cat{})
	} else {

		if idOfDeleting > 0 {
			if _, ok := catsDB[uint(idOfDeleting)]; ok {
				c.HTML(http.StatusOK, "catDelete.html", catsDB[uint(idOfDeleting)])
				delete(catsDB, uint(idOfDeleting))
			} else {
				c.HTML(http.StatusOK, "catDelete.html", Cat{})
			}
		}
	}

}
func readCat(c *gin.Context) {
	str := c.Query("id")
	if idOfReading, err := strconv.Atoi(str); err != nil {
		c.HTML(http.StatusOK, "catRead.html", Cat{})
	} else {

		if idOfReading > 0 {
			if _, ok := catsDB[uint(idOfReading)]; ok {
				c.HTML(http.StatusOK, "catRead.html", catsDB[uint(idOfReading)])
			} else {
				c.HTML(http.StatusOK, "catRead.html", Cat{})
			}
		}
	}
}
