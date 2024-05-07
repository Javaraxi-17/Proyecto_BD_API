package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetComentariosHandler(w http.ResponseWriter, r *http.Request) {
	var comentarios []models.Comentario
	fmt.Sprintln(comentarios)
	db.DB.Find(&comentarios)
	json.NewEncoder(w).Encode(&comentarios)

}

func GetComentarioHandler(w http.ResponseWriter, r *http.Request) {
	var comentario models.Comentario
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&comentario, params["id"])
	if comentario.ID_comentario == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&comentario)

}

func PostComentarioHandler(w http.ResponseWriter, r *http.Request) {
	var comentario models.Comentario

	json.NewDecoder(r.Body).Decode(&comentario)

	createdUser := db.DB.Create(&comentario)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&comentario)

}

func DeleteComentariosHandler(w http.ResponseWriter, r *http.Request) {
	var comentarios models.Comentario
	params := mux.Vars(r)

	db.DB.First(&comentarios, params["id"])
	if comentarios.ID_comentario == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&comentarios)
	w.WriteHeader(http.StatusAccepted)

}
