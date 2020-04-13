package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var client http.Client

func clientUser() {
	client = http.Client{
		Timeout: 10 * time.Second,
	}
}

func requestDeletingUserInDB(id int) (int, string) {
	idStr := strconv.Itoa(id)
	url := "http://localhost:3030/deleteUser?id=" + idStr
	resp, err := client.Get(url)
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}
	return resp.StatusCode, string(b)
}
func requestUpdatingUserInDB(u *User) (int, string, *User) {
	url := "http://localhost:3030/updateUser"
	marshalBytes, err := json.Marshal(u)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(marshalBytes))
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	if resp.StatusCode == http.StatusOK {
		user := User{}
		err := json.Unmarshal(b, &user)
		if err != nil {
			return http.StatusBadRequest, err.Error(), nil
		}
		return resp.StatusCode, "", &user
	}
	var msg string
	for _, value := range b {
		msg += string(value)
	}
	return resp.StatusCode, msg, nil
}
func requestCreatingUserInDB(u *User) (int, string, *User) {
	url := "http://localhost:3030/createUser"

	marshalBytes, err := json.Marshal(u)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(marshalBytes))
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	if resp.StatusCode == http.StatusOK {
		user := User{}
		err = json.Unmarshal(b, &user)
		if err != nil {
			return http.StatusBadRequest, err.Error(), nil
		}
		return resp.StatusCode, "", &user
	}
	var msg string
	for _, value := range b {
		msg += string(value)
	}
	return resp.StatusCode, msg, nil
}

func requestAutorization(u *User) (int, string, *User) {
	url := "http://localhost:3030/autorizationUser"

	marshalBytes, err := json.Marshal(u)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(marshalBytes))
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	if resp.StatusCode == http.StatusOK {
		var user User
		err := json.Unmarshal(b, &user)
		if err != nil {
			return http.StatusBadRequest, "wrong username or password", nil
		}
		return http.StatusOK, "", &user
	}
	var msg string
	for _, value := range b {
		msg += string(value)
	}
	return http.StatusBadRequest, msg, nil

}

func requestReading(id int) (int, string, *User) {
	idString := strconv.Itoa(id)
	url := "http://localhost:3030/readUser?id=" + idString
	resp, err := client.Get(url)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	if resp.StatusCode == http.StatusOK {
		var user User
		err := json.Unmarshal(b, &user)
		if err != nil {
			return http.StatusBadRequest, err.Error(), nil
		}
		return http.StatusOK, "", &user
	}
	var msg string
	for _, value := range b {
		msg += string(value)
	}
	return http.StatusBadRequest, msg, nil

}

func requestCatCreating(c *Cat, idOfUser int) (int, string, *User) {
	idStr := strconv.Itoa(idOfUser)
	url := "http://localhost:3030/createCat?id=" + idStr

	marshalBytes, err := json.Marshal(c)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(marshalBytes))
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	if resp.StatusCode == http.StatusOK {
		user := User{}
		err = json.Unmarshal(b, &user)
		if err != nil {
			return http.StatusBadRequest, err.Error(), nil
		}
		fmt.Println(user)

		return resp.StatusCode, "", &user
	}
	var msg string
	for _, value := range b {
		msg += string(value)
	}
	return resp.StatusCode, msg, nil

}

func requestActionAutorizationUser(cookie string) (int, string) {
	url := "http://localhost:3030/actionAutorizationUser"
	obj1 := obj{"cookie": cookie}
	marshalBytes, err := json.Marshal(obj1)
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(marshalBytes))
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}
	var msg string
	for _, value := range b {
		msg += string(value)
	}
	return resp.StatusCode, msg

}

func requestCatUpdating(c *Cat, idOfUser int) (int, string, *User) {
	idStr := strconv.Itoa(idOfUser)
	url := "http://localhost:3030/updateCat?id=" + idStr

	marshalBytes, err := json.Marshal(c)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(marshalBytes))
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	if resp.StatusCode == http.StatusOK {
		user := User{}
		err = json.Unmarshal(b, &user)
		if err != nil {
			return http.StatusBadRequest, err.Error(), nil
		}
		return resp.StatusCode, "", &user
	}
	var msg string
	for _, value := range b {
		msg += string(value)
	}
	return resp.StatusCode, msg, nil
}
func requestCatDeleting(idOfCat, idOfUser int) (int, string, *User) {
	idUserStr := strconv.Itoa(idOfUser)
	idCatStr := strconv.Itoa(idOfCat)
	url := "http://localhost:3030/deleteCat?idOfUser=" + idUserStr + "&idOfCat=" + idCatStr
	resp, err := client.Get(url)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil
	}
	if resp.StatusCode == http.StatusOK {
		user := User{}
		err = json.Unmarshal(b, &user)
		if err != nil {
			return http.StatusBadRequest, err.Error(), nil
		}
		return resp.StatusCode, "", &user
	}
	var msg string
	for _, value := range b {
		msg += string(value)
	}
	return resp.StatusCode, msg, nil

}
