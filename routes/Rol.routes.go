package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetRolesHandler(w http.ResponseWriter, r *http.Request) {
	var roles []models.Rol
	fmt.Sprintln(roles)
	db.DB.Find(&roles)
	json.NewEncoder(w).Encode(&roles)

}

func GetRolHandler(w http.ResponseWriter, r *http.Request) {
	var rol models.Rol
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&rol, params["id"])
	if rol.ID_rol == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&rol)

}

func PostRolHandler(w http.ResponseWriter, r *http.Request) {
	var rol models.Rol

	json.NewDecoder(r.Body).Decode(&rol)

	createdRol := db.DB.Create(&rol)
	err := createdRol.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&rol)

}

func DeleteRolesHandler(w http.ResponseWriter, r *http.Request) {
	var roles models.Rol
	params := mux.Vars(r)

	db.DB.First(&roles, params["id"])
	if roles.ID_rol == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&roles)
	w.WriteHeader(http.StatusAccepted)

}
