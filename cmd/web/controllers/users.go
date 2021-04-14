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
	database, err := db.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	repo := crud.NewRepository(database)

	users, err := repo.FindAll()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, users)
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
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	repo := crud.NewRepository(database)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := repo.Save(user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, res)
}
