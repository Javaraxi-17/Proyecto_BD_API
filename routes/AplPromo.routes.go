package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetAplicacionPromocionesHandler(w http.ResponseWriter, r *http.Request) {
	var AplicacionPromociones []models.AplicacionPromocion
	fmt.Sprintln(AplicacionPromociones)
	db.DB.Find(&AplicacionPromociones)
	json.NewEncoder(w).Encode(&AplicacionPromociones)

}

func GetAplicacionPromocionHandler(w http.ResponseWriter, r *http.Request) {
	var AplicacionPromocion models.AplicacionPromocion
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&AplicacionPromocion, params["id"])
	if AplicacionPromocion.ID_aplicacion_promocion == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&AplicacionPromocion)

}

func PostAplicacionPromocionHandler(w http.ResponseWriter, r *http.Request) {
	var AplicacionPromocion models.AplicacionPromocion

	json.NewDecoder(r.Body).Decode(&AplicacionPromocion)

	createdUser := db.DB.Create(&AplicacionPromocion)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&AplicacionPromocion)

}

func DeleteAplicacionPromocionesHandler(w http.ResponseWriter, r *http.Request) {
	var AplicacionPromociones models.AplicacionPromocion
	params := mux.Vars(r)

	db.DB.First(&AplicacionPromociones, params["id"])
	if AplicacionPromociones.ID_aplicacion_promocion == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&AplicacionPromociones)
	w.WriteHeader(http.StatusAccepted)

}
