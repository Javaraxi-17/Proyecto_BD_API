package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

// ciudad
func GetCiudadesHandler(w http.ResponseWriter, r *http.Request) {
	var Ciudades []models.Ciudad
	fmt.Sprintln(Ciudades)
	db.DB.Find(&Ciudades)
	json.NewEncoder(w).Encode(&Ciudades)

}

func GetCiudadHandler(w http.ResponseWriter, r *http.Request) {
	var Ciudad models.Ciudad
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&Ciudad, params["id"])
	if Ciudad.IDCiudad == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&Ciudad)

}

func PostCiudadHandler(w http.ResponseWriter, r *http.Request) {
	var Ciudad models.Ciudad

	json.NewDecoder(r.Body).Decode(&Ciudad)

	createdUser := db.DB.Create(&Ciudad)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Ciudad)

}

func DeleteCiudadesHandler(w http.ResponseWriter, r *http.Request) {
	var Ciudades models.Ciudad
	params := mux.Vars(r)

	db.DB.First(&Ciudades, params["id"])
	if Ciudades.IDCiudad == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&Ciudades)
	w.WriteHeader(http.StatusAccepted)

}

// barrio
func GetBarriosHandler(w http.ResponseWriter, r *http.Request) {
	var Barrios []models.Barrio
	fmt.Sprintln(Barrios)
	db.DB.Find(&Barrios)
	json.NewEncoder(w).Encode(&Barrios)

}

func GetBarrioHandler(w http.ResponseWriter, r *http.Request) {
	var Barrio models.Barrio
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&Barrio, params["id"])
	if Barrio.ID_barrio == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&Barrio)

}

func PostBarrioHandler(w http.ResponseWriter, r *http.Request) {
	var Barrio models.Barrio

	json.NewDecoder(r.Body).Decode(&Barrio)

	createdUser := db.DB.Create(&Barrio)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Barrio)

}

func DeleteBarriosHandler(w http.ResponseWriter, r *http.Request) {
	var Barrios models.Barrio
	params := mux.Vars(r)

	db.DB.First(&Barrios, params["id"])
	if Barrios.ID_barrio == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&Barrios)
	w.WriteHeader(http.StatusAccepted)

}
