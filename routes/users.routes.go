package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.Usuario
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	params := mux.Vars(r)

	fmt.Println(params["id"])

	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}

	json.NewEncoder(w).Encode(&user)

}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)

}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusAccepted)

}
