package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetCategoriaPropiedadesHandler(w http.ResponseWriter, r *http.Request) {
	var CategoriaPropiedades []models.CategoriaPropiedad
	fmt.Sprintln(CategoriaPropiedades)
	db.DB.Find(&CategoriaPropiedades)
	json.NewEncoder(w).Encode(&CategoriaPropiedades)

}

func GetCategoriaPropiedadHandler(w http.ResponseWriter, r *http.Request) {
	var CategoriaPropiedad models.CategoriaPropiedad
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&CategoriaPropiedad, params["id"])
	if CategoriaPropiedad.ID_categoria_propiedad == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&CategoriaPropiedad)

}

func PostCategoriaPropiedadHandler(w http.ResponseWriter, r *http.Request) {
	var CategoriaPropiedad models.CategoriaPropiedad

	json.NewDecoder(r.Body).Decode(&CategoriaPropiedad)

	createdUser := db.DB.Create(&CategoriaPropiedad)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&CategoriaPropiedad)

}

func DeleteCategoriaPropiedadesHandler(w http.ResponseWriter, r *http.Request) {
	var CategoriaPropiedad models.CategoriaPropiedad
	params := mux.Vars(r)

	db.DB.First(&CategoriaPropiedad, params["id"])
	if CategoriaPropiedad.ID_categoria_propiedad == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&CategoriaPropiedad)
	w.WriteHeader(http.StatusAccepted)

}
