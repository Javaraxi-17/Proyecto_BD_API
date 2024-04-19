package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetOwnersHandler(w http.ResponseWriter, r *http.Request) {
	var propietarios []models.Propietario
	fmt.Sprintln(propietarios)
	db.DB.Find(&propietarios)
	json.NewEncoder(w).Encode(&propietarios)

}

func GetOwnerHandler(w http.ResponseWriter, r *http.Request) {
	var propietario models.Propietario
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&propietario, params["id"])
	if propietario.IDPropietario == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&propietario)

}

func PostOwnerHandler(w http.ResponseWriter, r *http.Request) {
	var propietario models.Propietario

	json.NewDecoder(r.Body).Decode(&propietario)

	createdUser := db.DB.Create(&propietario)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&propietario)

}

func DeleteOwnerHandler(w http.ResponseWriter, r *http.Request) {
	var propietario models.Propietario
	params := mux.Vars(r)

	db.DB.First(&propietario, params["id"])
	if propietario.IDPropietario == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&propietario)
	w.WriteHeader(http.StatusAccepted)

}
