package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"wilbertopachecob/go_blog/cmd/web/db"
	"wilbertopachecob/go_blog/cmd/web/models"
	"wilbertopachecob/go_blog/cmd/web/repository/crud"
	"wilbertopachecob/go_blog/cmd/web/responses"
)

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
	database, err := db.Connect()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	repo := crud.NewRepository(database)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	res, err := repo.Save(user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responses.JSON(w, http.StatusCreated, res)
}
