package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetReservasHandler(w http.ResponseWriter, r *http.Request) {
	var Reservas []models.Reserva
	fmt.Sprintln(Reservas)
	db.DB.Find(&Reservas)
	json.NewEncoder(w).Encode(&Reservas)

}

func GetReservaHandler(w http.ResponseWriter, r *http.Request) {
	var Reserva models.Reserva
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&Reserva, params["id"])
	if Reserva.ID_reserva == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&Reserva)

}

func PostReservaHandler(w http.ResponseWriter, r *http.Request) {
	var Reserva models.Reserva

	json.NewDecoder(r.Body).Decode(&Reserva)

	createdUser := db.DB.Create(&Reserva)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Reserva)

}

func DeleteReservasHandler(w http.ResponseWriter, r *http.Request) {
	var Reservas models.Reserva
	params := mux.Vars(r)

	db.DB.First(&Reservas, params["id"])
	if Reservas.ID_reserva == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&Reservas)
	w.WriteHeader(http.StatusAccepted)

}

//reservahistorica

func GetReservasHistoricasHandler(w http.ResponseWriter, r *http.Request) {
	var ReservasHistoricas []models.ReservaHistorica
	fmt.Sprintln(ReservasHistoricas)
	db.DB.Find(&ReservasHistoricas)
	json.NewEncoder(w).Encode(&ReservasHistoricas)

}

func GetReservaHistoricaHandler(w http.ResponseWriter, r *http.Request) {
	var ReservaHistorica models.ReservaHistorica
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&ReservaHistorica, params["id"])
	if ReservaHistorica.ID_reserva_historica == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&ReservaHistorica)

}

func PostReservaHistoricaHandler(w http.ResponseWriter, r *http.Request) {
	var ReservaHistorica models.ReservaHistorica

	json.NewDecoder(r.Body).Decode(&ReservaHistorica)

	createdUser := db.DB.Create(&ReservaHistorica)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&ReservaHistorica)

}

func DeleteReservasHistoricasHandler(w http.ResponseWriter, r *http.Request) {
	var ReservasHistoricas models.ReservaHistorica
	params := mux.Vars(r)

	db.DB.First(&ReservasHistoricas, params["id"])
	if ReservasHistoricas.ID_reserva_historica == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&ReservasHistoricas)
	w.WriteHeader(http.StatusAccepted)

}
