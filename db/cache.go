package main

import (
	"strconv"
	"sync"
)

var userCache map[uint]User
var muxForUserCache sync.Mutex
var adminCache map[uint]Admin
var muxForAdminCache sync.Mutex
var userCookie map[string][2]string
var muxForUserCookie sync.Mutex
var adminCookie map[string][2]string
var muxForAdminCookie sync.Mutex

func defineUserCookie() {
	userCookie = make(map[string][2]string)
	for _, value := range userCache {
		temp := [2]string{value.Username, value.Password}
		userCookie[encryptCookie(value.Password)] = temp

	}
}
func defineUserCache() {
	userCache = make(map[uint]User)
	result := make([]User, 0)
	informationAboutUsers.Find(obj{"_id": obj{"$gt": 0}}).All(&result)
	for _, value := range result {
		userCache[value.ID] = value
	}
}

func defineAdminCookie() {
	adminCookie = make(map[string][2]string)
	for _, value := range adminCache {
		temp := [2]string{value.Username, value.Password}
		adminCookie[encryptCookie(value.Password)] = temp

	}
}
func defineAdminCache() {
	adminCache = make(map[uint]Admin)
	result := make([]Admin, 0)
	informationAboutAdmins.Find(obj{"_id": obj{"$gt": 0}}).All(&result)
	for _, value := range result {
		adminCache[value.ID] = value
	}
}
func encryptCookie(cookie string) (result string) {
	for i := 0; i < len(cookie); i++ {
		result += strconv.Itoa(int(cookie[i]))
	}
	return
}
