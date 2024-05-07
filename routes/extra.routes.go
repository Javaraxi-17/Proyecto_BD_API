package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

//Alertas

func GetAlertasHandler(w http.ResponseWriter, r *http.Request) {
	var Alertas []models.Alerta
	fmt.Sprintln(Alertas)
	db.DB.Find(&Alertas)
	json.NewEncoder(w).Encode(&Alertas)

}

func GetAlertaHandler(w http.ResponseWriter, r *http.Request) {
	var Alerta models.Alerta
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&Alerta, params["id"])
	if Alerta.ID_alerta == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&Alerta)

}

func PostAlertaHandler(w http.ResponseWriter, r *http.Request) {
	var Alerta models.Alerta

	json.NewDecoder(r.Body).Decode(&Alerta)

	createdUser := db.DB.Create(&Alerta)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Alerta)

}

func DeleteAlertasHandler(w http.ResponseWriter, r *http.Request) {
	var Alertas models.Alerta
	params := mux.Vars(r)

	db.DB.First(&Alertas, params["id"])
	if Alertas.ID_alerta == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&Alertas)
	w.WriteHeader(http.StatusAccepted)

}

//Configuracion

func GetConfiguracionesHandler(w http.ResponseWriter, r *http.Request) {
	var Configuraciones []models.Configuracion
	fmt.Sprintln(Configuraciones)
	db.DB.Find(&Configuraciones)
	json.NewEncoder(w).Encode(&Configuraciones)

}

func GetConfiguracionHandler(w http.ResponseWriter, r *http.Request) {
	var Configuracion models.Configuracion
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&Configuracion, params["id"])
	if Configuracion.ID_configuracion == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&Configuracion)

}

func PostConfiguracionHandler(w http.ResponseWriter, r *http.Request) {
	var Configuracion models.Configuracion

	json.NewDecoder(r.Body).Decode(&Configuracion)

	createdUser := db.DB.Create(&Configuracion)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Configuracion)

}

func DeleteConfiguracionesHandler(w http.ResponseWriter, r *http.Request) {
	var Configuraciones models.Configuracion
	params := mux.Vars(r)

	db.DB.First(&Configuraciones, params["id"])
	if Configuraciones.ID_configuracion == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&Configuraciones)
	w.WriteHeader(http.StatusAccepted)

}

//HistorialPrecio

func GetHistorialPreciosHandler(w http.ResponseWriter, r *http.Request) {
	var HistorialPrecios []models.HistorialPrecio
	fmt.Sprintln(HistorialPrecios)
	db.DB.Find(&HistorialPrecios)
	json.NewEncoder(w).Encode(&HistorialPrecios)

}

func GetHistorialPrecioHandler(w http.ResponseWriter, r *http.Request) {
	var HistorialPrecio models.HistorialPrecio
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&HistorialPrecio, params["id"])
	if HistorialPrecio.ID_historial_precio == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&HistorialPrecio)

}

func PostHistorialPrecioHandler(w http.ResponseWriter, r *http.Request) {
	var HistorialPrecio models.HistorialPrecio

	json.NewDecoder(r.Body).Decode(&HistorialPrecio)

	createdUser := db.DB.Create(&HistorialPrecio)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&HistorialPrecio)

}

func DeleteHistorialPreciosHandler(w http.ResponseWriter, r *http.Request) {
	var HistorialPrecios models.HistorialPrecio
	params := mux.Vars(r)

	db.DB.First(&HistorialPrecios, params["id"])
	if HistorialPrecios.ID_historial_precio == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&HistorialPrecios)
	w.WriteHeader(http.StatusAccepted)

}

//SolicitudAdministrativa

func GetSolicitudesAdministrativasHandler(w http.ResponseWriter, r *http.Request) {
	var SolicitudesAdministrativas []models.SolicitudAdministrativa
	fmt.Sprintln(SolicitudesAdministrativas)
	db.DB.Find(&SolicitudesAdministrativas)
	json.NewEncoder(w).Encode(&SolicitudesAdministrativas)

}

func GetSolicitudAdministrativaHandler(w http.ResponseWriter, r *http.Request) {
	var SolicitudAdministrativa models.SolicitudAdministrativa
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&SolicitudAdministrativa, params["id"])
	if SolicitudAdministrativa.ID_solicitud_administrativa == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&SolicitudAdministrativa)

}

func PostSolicitudAdministrativaHandler(w http.ResponseWriter, r *http.Request) {
	var SolicitudAdministrativa models.SolicitudAdministrativa

	json.NewDecoder(r.Body).Decode(&SolicitudAdministrativa)

	createdUser := db.DB.Create(&SolicitudAdministrativa)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&SolicitudAdministrativa)

}

func DeleteSolicitudesAdministrativasHandler(w http.ResponseWriter, r *http.Request) {
	var SolicitudesAdministrativas models.SolicitudAdministrativa
	params := mux.Vars(r)

	db.DB.First(&SolicitudesAdministrativas, params["id"])
	if SolicitudesAdministrativas.ID_solicitud_administrativa == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&SolicitudesAdministrativas)
	w.WriteHeader(http.StatusAccepted)

}
