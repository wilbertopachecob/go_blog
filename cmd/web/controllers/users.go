package controllers

import "net/http"

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete users"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create users"))
}
