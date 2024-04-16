package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetHomesHandler(w http.ResponseWriter, r *http.Request) {
	var casas []models.Casa
	fmt.Sprintln(casas)
	db.DB.Find(&casas)
	json.NewEncoder(w).Encode(&casas)

}

func GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	var casa models.Casa
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&casa, params["id"])
	if casa.ID_casa == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&casa)

}

func PostHomeHandler(w http.ResponseWriter, r *http.Request) {
	var casa models.Casa

	json.NewDecoder(r.Body).Decode(&casa)

	createdUser := db.DB.Create(&casa)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&casa)

}

func DeleteHomesHandler(w http.ResponseWriter, r *http.Request) {
	var casa models.Casa
	params := mux.Vars(r)

	db.DB.First(&casa, params["id"])
	if casa.ID_casa == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&casa)
	w.WriteHeader(http.StatusAccepted)

}
