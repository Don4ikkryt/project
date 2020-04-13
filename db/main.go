package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var idUser uint
var idAdmin uint

type obj map[string]interface{}

func main() {
	s, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	db := s.DB("users")
	informationAboutUsers = db.C("informationAboutUsers")
	informationAboutAdmins = db.C("informationAboutAdmins")
	uniqueUsername(informationAboutUsers)
	uniqueUsername(informationAboutAdmins)
	defineIDUser()
	defineIDAdmin()
	defineUserCache()
	defineAdminCache()
	defineUserCookie()
	defineAdminCookie()
	serverDB()
}

func uniqueUsername(c *mgo.Collection) {
	index := mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		fmt.Println(err)
	}
}

func defineIDUser() {
	var OBJ1 obj
	if err := informationAboutUsers.FindId(0).One(&OBJ1); err != nil {
		informationAboutUsers.Insert(obj{"_id": 0, "ID": 0})

	} else {
		idINT := OBJ1["ID"].(int)
		if idINT > 0 {
			idUser = uint(idINT)
		}
	}
}

func defineIDAdmin() {
	var OBJ1 obj
	if err := informationAboutAdmins.FindId(0).One(&OBJ1); err != nil {
		informationAboutAdmins.Insert(obj{"_id": 0, "ID": 0})

	} else {
		idINT := OBJ1["ID"].(int)
		if idINT > 0 {
			idAdmin = uint(idINT)
		}
	}
}
