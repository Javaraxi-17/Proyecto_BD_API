package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetPromocionesHandler(w http.ResponseWriter, r *http.Request) {
	var Promociones []models.Promocion
	fmt.Sprintln(Promociones)
	db.DB.Find(&Promociones)
	json.NewEncoder(w).Encode(&Promociones)

}

func GetPromocionHandler(w http.ResponseWriter, r *http.Request) {
	var Promocion models.Promocion
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&Promocion, params["id"])
	if Promocion.ID_promocion == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&Promocion)

}

func PostPromocionHandler(w http.ResponseWriter, r *http.Request) {
	var Promocion models.Promocion

	json.NewDecoder(r.Body).Decode(&Promocion)

	createdUser := db.DB.Create(&Promocion)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Promocion)

}

func DeletePromocionesHandler(w http.ResponseWriter, r *http.Request) {
	var Promociones models.Promocion
	params := mux.Vars(r)

	db.DB.First(&Promociones, params["id"])
	if Promociones.ID_promocion == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&Promociones)
	w.WriteHeader(http.StatusAccepted)

}
