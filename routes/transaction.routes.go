package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Javaraxi-17/Proyecto_BD_API/db"
	"github.com/Javaraxi-17/Proyecto_BD_API/models"
	"github.com/gorilla/mux"
)

func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transaccion
	fmt.Sprintln(transactions)
	db.DB.Find(&transactions)
	json.NewEncoder(w).Encode(&transactions)

}

func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var Transaction models.Transaccion
	params := mux.Vars(r)
	fmt.Println("estoy aqui")

	fmt.Println(params["id"])

	db.DB.First(&Transaction, params["id"])
	if Transaction.ID_transaccion == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Home not found"))
		return
	}

	json.NewEncoder(w).Encode(&Transaction)

}

func PostTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var Transaction models.Transaccion

	json.NewDecoder(r.Body).Decode(&Transaction)

	createdUser := db.DB.Create(&Transaction)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Transaction)

}

func DeleteTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	var Transaction models.Transaccion
	params := mux.Vars(r)

	db.DB.First(&Transaction, params["id"])
	if Transaction.ID_transaccion == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// db.DB.Delete(&user) //soft delete
	db.DB.Unscoped().Delete(&Transaction)
	w.WriteHeader(http.StatusAccepted)

}
