package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetPagosPendientesHandler(w http.ResponseWriter, r *http.Request) {
	var PagosPendientes []models.PagoPendiente
	fmt.Sprintln(PagosPendientes)
	db.DB.Find(&PagosPendientes)
	json.NewEncoder(w).Encode(&PagosPendientes)

}

func GetPagoPendienteHandler(w http.ResponseWriter, r *http.Request) {
	var PagoPendiente models.PagoPendiente
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&PagoPendiente, params["id"])
	if PagoPendiente.ID_pago_pendiente == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&PagoPendiente)

}

func PostPagoPendienteHandler(w http.ResponseWriter, r *http.Request) {
	var PagoPendiente models.PagoPendiente

	json.NewDecoder(r.Body).Decode(&PagoPendiente)

	createdUser := db.DB.Create(&PagoPendiente)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&PagoPendiente)

}

func DeletePagosPendientesHandler(w http.ResponseWriter, r *http.Request) {
	var PagosPendientes models.PagoPendiente
	params := mux.Vars(r)

	db.DB.First(&PagosPendientes, params["id"])
	if PagosPendientes.ID_pago_pendiente == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&PagosPendientes)
	w.WriteHeader(http.StatusAccepted)

}

//comision agente

func GetComisionAgentesHandler(w http.ResponseWriter, r *http.Request) {
	var ComisionAgentes []models.ComisionAgente
	fmt.Sprintln(ComisionAgentes)
	db.DB.Find(&ComisionAgentes)
	json.NewEncoder(w).Encode(&ComisionAgentes)

}

func GetComisionAgenteHandler(w http.ResponseWriter, r *http.Request) {
	var ComisionAgente models.ComisionAgente
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&ComisionAgente, params["id"])
	if ComisionAgente.ID_comision_agente == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&ComisionAgente)

}

func PostComisionAgenteHandler(w http.ResponseWriter, r *http.Request) {
	var ComisionAgente models.ComisionAgente

	json.NewDecoder(r.Body).Decode(&ComisionAgente)

	createdUser := db.DB.Create(&ComisionAgente)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&ComisionAgente)

}

func DeleteComisionAgentesHandler(w http.ResponseWriter, r *http.Request) {
	var ComisionAgentes models.ComisionAgente
	params := mux.Vars(r)

	db.DB.First(&ComisionAgentes, params["id"])
	if ComisionAgentes.ID_comision_agente == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&ComisionAgentes)
	w.WriteHeader(http.StatusAccepted)

}
