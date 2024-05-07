package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

// Bitacora accesos
func GetBitacoraAccesosHandler(w http.ResponseWriter, r *http.Request) {
	var BitacoraAccesos []models.BitacoraAcceso
	fmt.Sprintln(BitacoraAccesos)
	db.DB.Find(&BitacoraAccesos)
	json.NewEncoder(w).Encode(&BitacoraAccesos)

}

func GetBitacoraAccesoHandler(w http.ResponseWriter, r *http.Request) {
	var BitacoraAcceso models.BitacoraAcceso
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&BitacoraAcceso, params["id"])
	if BitacoraAcceso.ID_bitacora_acceso == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&BitacoraAcceso)

}

func PostBitacoraAccesoHandler(w http.ResponseWriter, r *http.Request) {
	var BitacoraAcceso models.BitacoraAcceso

	json.NewDecoder(r.Body).Decode(&BitacoraAcceso)

	createdUser := db.DB.Create(&BitacoraAcceso)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&BitacoraAcceso)

}

func DeleteBitacoraAccesosHandler(w http.ResponseWriter, r *http.Request) {
	var BitacoraAccesos models.BitacoraAcceso
	params := mux.Vars(r)

	db.DB.First(&BitacoraAccesos, params["id"])
	if BitacoraAccesos.ID_bitacora_acceso == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&BitacoraAccesos)
	w.WriteHeader(http.StatusAccepted)

}

// Bitacora actividad

func GetBitacoraActividadesHandler(w http.ResponseWriter, r *http.Request) {
	var BitacoraActividades []models.BitacoraActividad
	fmt.Sprintln(BitacoraActividades)
	db.DB.Find(&BitacoraActividades)
	json.NewEncoder(w).Encode(&BitacoraActividades)

}

func GetBitacoraActividadHandler(w http.ResponseWriter, r *http.Request) {
	var BitacoraActividad models.BitacoraActividad
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&BitacoraActividad, params["id"])
	if BitacoraActividad.ID_bitacora_actividad == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&BitacoraActividad)

}

func PostBitacoraActividadHandler(w http.ResponseWriter, r *http.Request) {
	var BitacoraActividad models.BitacoraActividad

	json.NewDecoder(r.Body).Decode(&BitacoraActividad)

	createdUser := db.DB.Create(&BitacoraActividad)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&BitacoraActividad)

}

func DeleteBitacoraActividadesHandler(w http.ResponseWriter, r *http.Request) {
	var BitacoraActividades models.BitacoraActividad
	params := mux.Vars(r)

	db.DB.First(&BitacoraActividades, params["id"])
	if BitacoraActividades.ID_bitacora_actividad == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&BitacoraActividades)
	w.WriteHeader(http.StatusAccepted)

}

//Auditoria

func GetAuditoriasHandler(w http.ResponseWriter, r *http.Request) {
	var Auditorias []models.Auditoria
	fmt.Sprintln(Auditorias)
	db.DB.Find(&Auditorias)
	json.NewEncoder(w).Encode(&Auditorias)

}

func GetAuditoriaHandler(w http.ResponseWriter, r *http.Request) {
	var Auditoria models.Auditoria
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&Auditoria, params["id"])
	if Auditoria.ID_auditoria == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&Auditoria)

}

func PostAuditoriaHandler(w http.ResponseWriter, r *http.Request) {
	var Auditoria models.Auditoria

	json.NewDecoder(r.Body).Decode(&Auditoria)

	createdUser := db.DB.Create(&Auditoria)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Auditoria)

}

func DeleteAuditoriasHandler(w http.ResponseWriter, r *http.Request) {
	var Auditorias models.Auditoria
	params := mux.Vars(r)

	db.DB.First(&Auditorias, params["id"])
	if Auditorias.ID_auditoria == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&Auditorias)
	w.WriteHeader(http.StatusAccepted)

}
