package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetPaisesHandler(w http.ResponseWriter, r *http.Request) {
	var Paises []models.Pais
	fmt.Sprintln(Paises)
	db.DB.Find(&Paises)
	json.NewEncoder(w).Encode(&Paises)

}

func GetPaisHandler(w http.ResponseWriter, r *http.Request) {
	var Pais models.Pais
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&Pais, params["id"])
	if Pais.ID_pais == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&Pais)

}

func PostPaisHandler(w http.ResponseWriter, r *http.Request) {
	var Pais models.Pais

	json.NewDecoder(r.Body).Decode(&Pais)

	createdUser := db.DB.Create(&Pais)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Pais)

}

func DeletePaisesHandler(w http.ResponseWriter, r *http.Request) {
	var Paises models.Pais
	params := mux.Vars(r)

	db.DB.First(&Paises, params["id"])
	if Paises.ID_pais == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&Paises)
	w.WriteHeader(http.StatusAccepted)

}
