package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetAgentesInmobiliariosHandler(w http.ResponseWriter, r *http.Request) {
	var AgentesInmobiliarios []models.AgenteInmobiliario
	fmt.Sprintln(AgentesInmobiliarios)
	db.DB.Find(&AgentesInmobiliarios)
	json.NewEncoder(w).Encode(&AgentesInmobiliarios)

}

func GetAgenteInmobiliarioHandler(w http.ResponseWriter, r *http.Request) {
	var AgenteInmobiliario models.AgenteInmobiliario
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&AgenteInmobiliario, params["id"])
	if AgenteInmobiliario.ID_agente_inmobiliario == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&AgenteInmobiliario)

}

func PostAgenteInmobiliarioHandler(w http.ResponseWriter, r *http.Request) {
	var AgenteInmobiliario models.AgenteInmobiliario

	json.NewDecoder(r.Body).Decode(&AgenteInmobiliario)

	createdUser := db.DB.Create(&AgenteInmobiliario)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&AgenteInmobiliario)

}

func DeleteAgentesInmobiliariosHandler(w http.ResponseWriter, r *http.Request) {
	var AgentesInmobiliarios models.AgenteInmobiliario
	params := mux.Vars(r)

	db.DB.First(&AgentesInmobiliarios, params["id"])
	if AgentesInmobiliarios.ID_agente_inmobiliario == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&AgentesInmobiliarios)
	w.WriteHeader(http.StatusAccepted)

}
