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

// seguimiento casa
func GetSeguimientoCasasHandler(w http.ResponseWriter, r *http.Request) {
	var SeguimientoCasas []models.SeguimientoCasa
	fmt.Sprintln(SeguimientoCasas)
	db.DB.Find(&SeguimientoCasas)
	json.NewEncoder(w).Encode(&SeguimientoCasas)

}

func GetSeguimientoCasaHandler(w http.ResponseWriter, r *http.Request) {
	var SeguimientoCasa models.SeguimientoCasa
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&SeguimientoCasa, params["id"])
	if SeguimientoCasa.ID_seguimiento_casa == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&SeguimientoCasa)

}

func PostSeguimientoCasaHandler(w http.ResponseWriter, r *http.Request) {
	var SeguimientoCasa models.SeguimientoCasa

	json.NewDecoder(r.Body).Decode(&SeguimientoCasa)

	createdUser := db.DB.Create(&SeguimientoCasa)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&SeguimientoCasa)

}

func DeleteSeguimientoCasasHandler(w http.ResponseWriter, r *http.Request) {
	var SeguimientoCasas models.SeguimientoCasa
	params := mux.Vars(r)

	db.DB.First(&SeguimientoCasas, params["id"])
	if SeguimientoCasas.ID_seguimiento_casa == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&SeguimientoCasas)
	w.WriteHeader(http.StatusAccepted)

}
