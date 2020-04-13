package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var client http.Client

func clientAdmin() {
	client = http.Client{
		Timeout: 10 * time.Second,
	}
}

func requestWriteMessage(id int, msg string) (int, string) {
	url := "http://localhost:3030/message"
	marshalBytes, err := json.Marshal(obj{"message": msg, "id": id})
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
	var messageResp string
	for _, value := range b {
		messageResp += string(value)
	}
	return resp.StatusCode, messageResp

}

func requestReadingUsers() (int, string, []User) {
	url := "http://localhost:3030/readUsers"
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
		users := make([]User, 0)
		err = json.Unmarshal(b, &users)
		if err != nil {
			return http.StatusBadRequest, err.Error(), nil
		}
		return resp.StatusCode, "", users
	}
	var msg string
	for _, value := range b {
		msg += string(value)
	}
	return resp.StatusCode, msg, nil

}

func requestAutorization(a *Admin) (int, string, *Admin) {
	url := "http://localhost:3030/autorizationAdmin"

	marshalBytes, err := json.Marshal(a)
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
		var admin Admin
		err := json.Unmarshal(b, &admin)
		if err != nil {
			return http.StatusBadRequest, "wrong username or password", nil
		}
		return http.StatusOK, "", &admin
	}
	var msg string
	for _, value := range b {
		msg += string(value)
	}
	return http.StatusBadRequest, msg, nil

}
func requestCreatingAdminInDB(a *Admin) (int, string, *Admin) {
	url := "http://localhost:3030/createAdmin"

	marshalBytes, err := json.Marshal(a)
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
		admin := Admin{}
		err = json.Unmarshal(b, &admin)
		if err != nil {
			return http.StatusBadRequest, err.Error(), nil
		}
		return resp.StatusCode, "", &admin
	}
	var msg string
	for _, value := range b {
		msg += string(value)
	}
	return resp.StatusCode, msg, nil
}
func requestActionAutorizationAdmin(cookie string) (int, string) {
	url := "http://localhost:3030/actionAutorizationAdmin"
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
func requestReading(id int) (int, string, *User) {
	idString := strconv.Itoa(id)
	url := "http://localhost:3030/readAdmin?id=" + idString
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
